package mappers

import (
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/models/svg"
	"github.com/MagnusHLund/Github-profile-stats-tracker/internal/utils"
)

type QueryParameterMapper struct {
	queryParameters *utils.QueryParameters
	helperUtils     *utils.HelperUtils
}

func NewQueryParameterMapper() *QueryParameterMapper {
	return &QueryParameterMapper{}
}

func (q *QueryParameterMapper) MapQueryParametersToSpecifiedSvgType(queryParameters *utils.QueryParameters) svg.IAbstractSVG {
	if queryParameters.SVGType == "visitor_counter" {
		return q.MapQueryParametersToVisitorCounterSVG(queryParameters)
	}

	return nil
}

func (q *QueryParameterMapper) MapQueryParametersToVisitorCounterSVG(queryParameters *utils.QueryParameters) svg.VisitorCounterSVG {
	return svg.VisitorCounterSVG{
		AbstractSVG: q.mapQueryParametersToAbstractSVG(queryParameters),
		Text:        svg.Text{Content: queryParameters.Text, Fill: queryParameters.Fill},
	}

}

func (q *QueryParameterMapper) mapQueryParametersToAbstractSVG(queryParameters *utils.QueryParameters) svg.AbstractSVG {
	return svg.AbstractSVG{
		Rect: struct {
			Width  string `xml:"width,attr"`
			Height string `xml:"height,attr"`
		}{
			Width:  q.helperUtils.DefaultIfEmpty(queryParameters.Width, "220"),
			Height: q.helperUtils.DefaultIfEmpty(queryParameters.Height, "50"),
		},
		Defs: struct {
			LinearGradient struct {
				Stops []struct {
					Offset string `xml:"offset,attr"`
					Style  string `xml:"style,attr"`
				} `xml:"stop"`
			} `xml:"linearGradient"`
		}{
			LinearGradient: struct {
				Stops []struct {	
					Offset string `xml:"offset,attr"`
					Style  string `xml:"style,attr"`
				}{
					{Offset: "0%", Style: "stop-color:" + queryParameters.StartColor},
					{Offset: "100%", Style: "stop-color:" + queryParameters.EndColor},
					
				},
			},
		},
	}
}
