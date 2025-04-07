package services

import (
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

func (vs *VisitorService) CreateVisitor(profilePage string, ipAddress string) {
	page, err := vs.PageService.CreatePageIfNotExists(profilePage)
	if err != nil {
		return
	}

	hashedIPAdress := vs.HashingService.Hash(ipAddress)
	vs.VisitorRepository.CreateVisitorIfNotExistsForPage(page.PageId, hashedIPAdress)
}
