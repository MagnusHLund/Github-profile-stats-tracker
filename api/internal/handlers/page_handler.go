package handlers

import (
	"net/http"

	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/services"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/utils"
)

type PageHandler struct {
	RequestUtils   *utils.RequestUtils
	PageService    *services.PageService
	VisitorService *services.VisitorService
}

func NewPageHandler(pageService *services.PageService, requestUtils *utils.RequestUtils, visitorService *services.VisitorService) *PageHandler {
	return &PageHandler{PageService: pageService, RequestUtils: requestUtils, VisitorService: visitorService}
}

func (ph *PageHandler) GetVisitorsForPage(w http.ResponseWriter, r *http.Request) {
	const ip = "" //ph.RequestUtils.GetIPAddress(r)
	ph.VisitorService.CreateVisitor(ip)
	ph.PageService.GetVisitorsForPage(w, r)
}
