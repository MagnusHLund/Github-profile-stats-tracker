package services

import (
	"github.com/MagnusHLund/VisitorCounter/internal/repositories"
)

type VisitorService struct {
	VisitorRepository *repositories.VisitorRepository
	HashingService    *HashingService
}

func NewVisitorService(visitorRepository *repositories.VisitorRepository, hashingService *HashingService) *VisitorService {
	return &VisitorService{VisitorRepository: visitorRepository, HashingService: hashingService}
}

func (vs *VisitorService) CreateVisitor(ipAddress string) {
	hashedIPAdress := vs.HashingService.Hash(ipAddress)
	vs.VisitorRepository.CreateVisitor(hashedIPAdress)
}
