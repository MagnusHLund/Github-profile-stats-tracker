package utils

import (
	"encoding/xml"
	"os"
	"path/filepath"
)

type SvgUtils struct{}

func NewSvgUtils() *SvgUtils {
	return &SvgUtils{}
}

func (su *SvgUtils) LoadSVGFromFile(svgType string) (string, error) {
	filePath := filepath.Join("assets", "images", svgType+".svg")
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var svg string
	err = xml.Unmarshal(data, &svg)
	if err != nil {
		return "", err
	}

	return svg, nil
}
