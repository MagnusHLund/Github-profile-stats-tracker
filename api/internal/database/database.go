package database

import (
	"fmt"
	"log"

	"github.com/MagnusHLund/VisitorCounter/internal/config"
	"github.com/MagnusHLund/VisitorCounter/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	if err := createDatabaseIfNotExists(cfg); err != nil {
		return nil, err
	}

	db, err := connectToDatabase(cfg)
	if err != nil {
		return nil, err
	}

	if err := migrateDatabase(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createDatabaseIfNotExists(cfg *config.Config) error {
	dsnWithoutDB := getConnectionString(cfg, false)

	tempDB, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{})
	if err != nil {
		return err
	}

	createDBQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.DatabaseConfig.DBName)
	if err := tempDB.Exec(createDBQuery).Error; err != nil {
		return err
	}

	return nil
}

func connectToDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := getConnectionString(cfg, true)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Visitor{},
		&models.Page{},
	); err != nil {
		log.Println("Failed to migrate database: ", err)
		return err
	}

	return nil
}

func getConnectionString(cfg *config.Config, databaseExists bool) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		cfg.DatabaseConfig.DBUser,
		cfg.DatabaseConfig.DBPass,
		cfg.DatabaseConfig.DBHost,
		cfg.DatabaseConfig.DBPort,
	)

	if databaseExists {
		dsn += cfg.DatabaseConfig.DBName
	}

	dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}
