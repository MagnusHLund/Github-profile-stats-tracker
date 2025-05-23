package repositories

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models/orm"
	"gorm.io/gorm"
)

type VisitorRepository struct {
	DB *gorm.DB
}

func NewVisitorRepository(db *gorm.DB) *VisitorRepository {
	return &VisitorRepository{DB: db}
}

func (r *VisitorRepository) CreateVisitorIfNotExistsForPage(pageId uint, hashedIPAdress string) error {
	visitor := &orm.Visitor{PageId: pageId, HashedIpAddress: hashedIPAdress}
	err := r.DB.Create(visitor).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *VisitorRepository) GetVisitorCountByPage(pageId uint) (uint, error) {
	visitors := []*orm.Visitor{}
	err := r.DB.Where("page_id = ?", pageId).Find(&visitors).Error

	if err != nil {
		return 0, err
	}

	visitorCount := uint(len(visitors))
	return visitorCount, nil
}
