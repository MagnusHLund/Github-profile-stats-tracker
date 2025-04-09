package svg

type VisitorCounterSVG struct {
	AbstractSVG
	Text struct {
		Content string `xml:",chardata"`
		Fill    string `xml:"fill,attr"`
	} `xml:"text"`
}
