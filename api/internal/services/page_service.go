package services

import (
	"encoding/json"
	"net/http"

	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/repositories"
)

type PageService struct {
	PageRepository *repositories.PageRepository
}

func NewPageService(PageRepository *repositories.PageRepository) *PageService {
	return &PageService{PageRepository: PageRepository}
}

func (ps *PageService) GetVisitorsForPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "List of pages"})
}
