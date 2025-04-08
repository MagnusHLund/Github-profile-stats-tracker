package utils

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"regexp"

	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models"
)

type SvgUtils struct{}

func NewSvgUtils() *SvgUtils {
	return &SvgUtils{}
}

func (su *SvgUtils) LoadSVGFromFile(statsType string) (*models.SVG, error) {
	filePath := filepath.Join("assets", "images", statsType+".svg")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the SVG
	var svg models.SVG
	err = xml.Unmarshal(data, &svg)
	if err != nil {
		return nil, err
	}
	return &svg, nil
}

func (su *SvgUtils) ModifySVG(svg *models.SVG, visitorCount uint, width string, height string, gradientStart string, gradientEnd string) (*models.SVG, error) {
	// TODO
	//setSVGGradient(svg, gradientStart, gradientEnd)

	return svg, nil
}

func setSVGGradient(svg *models.SVG, gradientStart *string, gradientEnd *string) {

	if gradientStart == nil || gradientEnd == nil {
		return
	}

	if !isValidHexColor(*gradientStart) || !isValidHexColor(*gradientEnd) {
		return
	}

	if len(svg.Defs.LinearGradient.Stops) >= 2 {
		svg.Defs.LinearGradient.Stops[0].Style = "stop-color:" + *gradientStart + ";stop-opacity:1"
		svg.Defs.LinearGradient.Stops[1].Style = "stop-color:" + *gradientEnd + ";stop-opacity:1"
	}
}

func isValidHexColor(color string) bool {
	const hexColorPattern = "^#([a-fA-F0-9]{3}|[a-fA-F0-9]{6})$"

	matched, _ := regexp.MatchString(hexColorPattern, color)
	return matched
}
