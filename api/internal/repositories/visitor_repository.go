package repositories

import "gorm.io/gorm"

type VisitorRepository struct {
	DB *gorm.DB
}

func NewVisitorRepository(db *gorm.DB) *VisitorRepository {
	return &VisitorRepository{DB: db}
}
