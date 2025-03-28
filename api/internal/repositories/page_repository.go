package repositories

import "gorm.io/gorm"

type PageRepository struct {
	DB *gorm.DB
}

func NewPageRepository(db *gorm.DB) *PageRepository {
	return &PageRepository{DB: db}
}
