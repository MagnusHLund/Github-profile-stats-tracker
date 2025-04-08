package svg

import "encoding/xml"

type AbstractSVG struct {
	XMLName xml.Name `xml:"svg"`
	Rect    struct {
		Width  string `xml:"width,attr"`
		Height string `xml:"height,attr"`
	} `xml:"rect"`
}
