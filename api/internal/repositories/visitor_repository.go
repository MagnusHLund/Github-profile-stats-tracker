package repositories

import (
	"github.com/MagnusHLund/VisitorCounter/internal/models"
	"gorm.io/gorm"
)

type VisitorRepository struct {
	DB *gorm.DB
}

func NewVisitorRepository(db *gorm.DB) *VisitorRepository {
	return &VisitorRepository{DB: db}
}

func (r *VisitorRepository) GetVisitorCountByPage(page *models.Page) (uint, error) {
	visitors := []*models.Visitor{}
	err := r.DB.Where("page_id = ?", page.PageId).Find(&visitors).Error

	if err != nil {
		return 0, err
	}

	return uint(len(visitors)), nil
}
