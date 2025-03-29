package services

import (
	"encoding/json"
	"net/http"

	"github.com/MagnusHLund/VisitorCounter/internal/repositories"
)

type VisitorService struct {
	VisitorRepository *repositories.VisitorRepository
}

func NewVisitorService(visitorRepository *repositories.VisitorRepository) *VisitorService {
	return &VisitorService{VisitorRepository: visitorRepository}
}

func (vs *VisitorService) CreateVisitorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Visitor created"})
}
