package handlers

import (
	"net/http"

	"github.com/MagnusHLund/VisitorCounter/internal/services"
)

type PageHandler struct {
	PageService *services.PageService
}

func NewPageHandler(pageService *services.PageService) *PageHandler {
	return &PageHandler{PageService: pageService}
}

func (h *PageHandler) GetVisitorsForPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Handle GET GET /page"))
}
