package handlers

import (
	"net/http"

	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/mappers"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/services"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/utils"
)

type PageHandler struct {
	RequestUtils         *utils.RequestUtils
	PageService          *services.PageService
	VisitorService       *services.VisitorService
	QueryParameterMapper *mappers.QueryParameterMapper
}

func NewPageHandler(pageService *services.PageService, requestUtils *utils.RequestUtils, visitorService *services.VisitorService) *PageHandler {
	return &PageHandler{PageService: pageService, RequestUtils: requestUtils, VisitorService: visitorService}
}

func (ph *PageHandler) GetVisitorsForPage(w http.ResponseWriter, r *http.Request) {
	profilePageOwner := ph.RequestUtils.GetPageOwnerGitUsername(r)
	ip := ph.RequestUtils.GetIPAddress(r)
	queryParameters := ph.RequestUtils.ParseQueryParameters(r)
	
	profilePage := ph.PageService.CreatePageIfNotExists(profilePageOwner)
	ph.VisitorService.CreateVisitor(profilePage, ip)

	svgOptions := ph.QueryParameterMapper.MapQueryParametersToAbstractSVG(queryParameters, "VisitorCount")
	ph.PageService.GetVisitorsForPage(profilePage, svgOptions)
}
