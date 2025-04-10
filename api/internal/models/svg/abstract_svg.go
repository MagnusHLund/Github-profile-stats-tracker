package svg

type IAbstractSVG interface {
	GetSVGType() string
}

type AbstractSVG struct {
	SVGType   string
	XMLName   string   `xml:"svg"`
	Width     string   `xml:"width,attr"`
	Height    string   `xml:"height,attr"`
	Gradient  Gradient `xml:"defs>linearGradient"`
	TextColor string   `xml:"style,attr"`
}

type Gradient struct {
	Stops []Stop `xml:"stop"`
}

type Stop struct {
	Offset    string `xml:"offset,attr"`
	StopColor string `xml:"style,attr"`
}

func (a *AbstractSVG) GetSVGType() string {
	return a.SVGType
}
