package utils

import (
	"encoding/xml"
	"path/filepath"
)

type SVGUtils struct {
	fileUtils *FileUtils
}

func NewSvgUtils(fileUtils *FileUtils) *SVGUtils {
	return &SVGUtils{fileUtils: fileUtils}
}

func (su *SVGUtils) LoadSVGFromFile(svgType string) (string, error) {
	filePathFromRoot := filepath.Join("internal", "assets", "svgs", svgType+".svg")
	data, err := su.fileUtils.ReadFile(filePathFromRoot)
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
