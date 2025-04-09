package svg

import "encoding/xml"

type IAbstractSVG interface {
}

type AbstractSVG struct {
	XMLName xml.Name `xml:"svg"`
	Rect    struct {
		Width  string `xml:"width,attr"`
		Height string `xml:"height,attr"`
	} `xml:"rect"`
	Defs struct {
		LinearGradient struct {
			Stops []struct {
				Offset string `xml:"offset,attr"`
				Style  string `xml:"style,attr"`
			} `xml:"stop"`
		} `xml:"linearGradient"`
	} `xml:"defs"`
}
