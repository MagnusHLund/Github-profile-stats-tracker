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
	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort)
	tempDB, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{})
	if err != nil {
		return err
	}

	createDBQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.DBName)
	if err := tempDB.Exec(createDBQuery).Error; err != nil {
		return err
	}

	return nil
}

func connectToDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

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
