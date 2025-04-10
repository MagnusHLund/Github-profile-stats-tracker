package svg

type VisitorCounterSVG struct {
	AbstractSVG
	VisitorCount string `xml:"chardata"`
}
