package services

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models/orm"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/repositories"
)

type VisitorService struct {
	VisitorRepository *repositories.VisitorRepository
	PageService       *PageService
	HashingService    *HashingService
}

func NewVisitorService(visitorRepository *repositories.VisitorRepository, hashingService *HashingService) *VisitorService {
	return &VisitorService{VisitorRepository: visitorRepository, HashingService: hashingService}
}

func (vs *VisitorService) CreateVisitor(profilePage *orm.Page, ipAddress string) {

	hashedIPAdress := vs.HashingService.Hash(ipAddress)
	vs.VisitorRepository.CreateVisitorIfNotExistsForPage(profilePage.PageId, hashedIPAdress)
}

func (vs *VisitorService) GetVisitorCountForPage(profilePage *orm.Page) uint {
	visitorCount, err := vs.VisitorRepository.GetVisitorCountByPage(profilePage.PageId)

	if err != nil {
		return 0
	}

	return visitorCount
}
