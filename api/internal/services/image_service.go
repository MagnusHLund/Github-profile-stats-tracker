package services

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/utils"
)

type ImageService struct {
	svgUtils *utils.SvgUtils
}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (is *ImageService) GenerateImageForPage(statsType string) {
	svgImage, err := s.svgUtils.LoadSVGFromFile(statsType)
	if err != nil {

		return
	}

	// TODO
	//is.svgUtils.ModifySVG(svgImage)
}
