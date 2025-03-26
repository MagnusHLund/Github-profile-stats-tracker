package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

type PageHandler struct {
	DB *gorm.DB
}

func NewPageHandler(db *gorm.DB) *PageHandler {
	return &PageHandler{DB: db}
}

func (h *PageHandler) CreatePage(w http.ResponseWriter, r *http.Request) {
	// Business logic here
	w.Write([]byte("Handle POST /page"))
}
