package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

type VisitorHandler struct {
	DB *gorm.DB
}

func NewVisitorHandler(db *gorm.DB) *VisitorHandler {
	return &VisitorHandler{DB: db}
}

func (h *VisitorHandler) GetVisitorsForPage(w http.ResponseWriter, r *http.Request) {
	// Business logic here
	w.Write([]byte("Handle GET /visitors"))
}

func (h *VisitorHandler) CreateVisitor(w http.ResponseWriter, r *http.Request) {
	
	w.Write([]byte("Handle POST /visitor"))
}
