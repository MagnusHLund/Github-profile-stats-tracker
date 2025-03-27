package services

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type VisitorService struct {
	db *gorm.DB
}

func NewVisitorService(db *gorm.DB) *VisitorService {
	return &VisitorService{db: db}
}

func (vs *VisitorService) GetVisitorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "List of visitors"})
}

func (vs *VisitorService) CreateVisitorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Visitor created"})
}
