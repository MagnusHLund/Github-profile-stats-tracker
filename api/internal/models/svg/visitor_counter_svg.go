package svg

type VisitorCounterSVG struct {
	AbstractSVG
	Defs struct {
		LinearGradient struct {
			Stops []struct {
				Offset string `xml:"offset,attr"`
				Style  string `xml:"style,attr"`
			} `xml:"stop"`
		} `xml:"linearGradient"`
	} `xml:"defs"`
	Text struct {
		Content string `xml:",chardata"`
		Fill    string `xml:"fill,attr"`
	} `xml:"text"`
}
