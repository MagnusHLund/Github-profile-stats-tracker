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

func (q *QueryParameterMapper) MapQueryParametersToAbstractSVG(queryParameters *utils.QueryParameters, svgType string) svg.AbstractSVG {
	return svg.AbstractSVG{
		SVGType:   svgType,
		Width:     q.helperUtils.DefaultIfEmpty(queryParameters.Width, "100"),
		Height:    q.helperUtils.DefaultIfEmpty(queryParameters.Height, "100"),
		TextColor: q.helperUtils.DefaultIfEmpty(queryParameters.TextColor, "defaultColor"),
		Gradient: svg.Gradient{
			Stops: []svg.Stop{
				{
					Offset:    "0%",
					StopColor: q.helperUtils.DefaultIfEmpty(queryParameters.GradientStart, "defaultColor"),
				},
				{
					Offset:    "100%",
					StopColor: q.helperUtils.DefaultIfEmpty(queryParameters.GradientEnd, "defaultColor"),
				},
			},
		},
	}
}
