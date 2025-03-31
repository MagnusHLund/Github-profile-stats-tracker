package repositories

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models"
	"gorm.io/gorm"
)

type VisitorRepository struct {
	DB *gorm.DB
}

func NewVisitorRepository(db *gorm.DB) *VisitorRepository {
	return &VisitorRepository{DB: db}
}

func (r *VisitorRepository) CreateVisitor(hashedIPAdress string) error {
	visitor := &models.Visitor{HashedIpAddress: hashedIPAdress}
	err := r.DB.Create(visitor).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *VisitorRepository) GetVisitorCountByPage(page *models.Page) (uint, error) {
	visitors := []*models.Visitor{}
	err := r.DB.Where("page_id = ?", page.PageId).Find(&visitors).Error

	if err != nil {
		return 0, err
	}

	return uint(len(visitors)), nil
}
