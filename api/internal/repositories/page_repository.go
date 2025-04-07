package repositories

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models"
	"gorm.io/gorm"
)

type PageRepository struct {
	DB *gorm.DB
}

func NewPageRepository(db *gorm.DB) *PageRepository {
	return &PageRepository{DB: db}
}

func (r *PageRepository) GetPageByGitUsername(url string) (*models.Page, error) {
	var page models.Page
	if err := r.DB.Where("url = ?", url).First(&page).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (r *PageRepository) CreatePage(page *models.Page) error {
	return r.DB.Create(page).Error
}
