package repositories

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models/orm"
	"gorm.io/gorm"
)

type PageRepository struct {
	DB *gorm.DB
}

func NewPageRepository(db *gorm.DB) *PageRepository {
	return &PageRepository{DB: db}
}

func (r *PageRepository) GetPageByGitUsername(gitUsername string) (*orm.Page, error) {
	var page orm.Page
	if err := r.DB.Where("page_owner_git_username = ?", gitUsername).First(&page).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (r *PageRepository) CreatePage(page *orm.Page) error {
	return r.DB.Create(page).Error
}
