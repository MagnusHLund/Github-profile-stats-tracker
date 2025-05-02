package services

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models/svg"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/utils"
)

type ImageService struct {
	svgUtils *utils.SVGUtils
}

func NewImageService(svgUtils *utils.SVGUtils) *ImageService {
	return &ImageService{svgUtils: svgUtils}
}

func (is *ImageService) GenerateVisitorCountImage(visitorCounterSVG svg.VisitorCounterSVG) []byte {
	svgImage, err := is.svgUtils.LoadSVGFromFile(visitorCounterSVG.SVGType)
	if err != nil {
		return nil
	}

	

	// Convert modifiedSVG to byte array and return
	return []byte(svgImage)
}
