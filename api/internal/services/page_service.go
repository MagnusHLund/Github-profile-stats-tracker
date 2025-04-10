package services

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models/orm"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models/svg"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/repositories"
)

type PageService struct {
	PageRepository *repositories.PageRepository
	VisitorService *VisitorService
	ImageService   *ImageService
}

func NewPageService(PageRepository *repositories.PageRepository, visitorService *VisitorService, imageService *ImageService) *PageService {
	return &PageService{PageRepository: PageRepository, VisitorService: visitorService, ImageService: imageService}
}

func (ps *PageService) GetVisitorsForPage(profilePage *orm.Page, svgOptions svg.AbstractSVG) {
	visitorCount := ps.VisitorService.GetVisitorCountForPage(profilePage)
	visitorCountSVG := svg.VisitorCounterSVG{AbstractSVG: svgOptions, VisitorCount: string(visitorCount)}

	image := ps.ImageService.GenerateVisitorCountImage(visitorCountSVG)
	println(string(image))
}

func (ps *PageService) CreatePageIfNotExists(profilePage string) *orm.Page {
	page, err := ps.PageRepository.GetPageByGitUsername(profilePage)

	if page == nil {
		page = &orm.Page{PageId: 0, PageOwnerGitUsername: profilePage}
		err = ps.PageRepository.CreatePage(page)
		if err != nil {
			return nil
		}
	}

	return page
}
