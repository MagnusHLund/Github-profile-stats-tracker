package services

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/repositories"
)

type PageService struct {
	PageRepository *repositories.PageRepository
}

func NewPageService(PageRepository *repositories.PageRepository) *PageService {
	return &PageService{PageRepository: PageRepository}
}

func (ps *PageService) GetVisitorsForPage(profilePage string)  {

}

func (ps *PageService) CreatePageIfNotExists(profilePage string) (*models.Page, error) {
	page, err := ps.PageRepository.GetPageByGitUsername(profilePage)
	if err != nil {
		return nil, err
	}

	if page == nil {
		page = &models.Page{PageId: 0, PageOwnerGitUsername: profilePage}
		err = ps.PageRepository.CreatePage(page)
		if err != nil {
			return nil, err
		}
	}

	return page, nil
}
