package utils

import (
	"net/http"
	"net/url"
	"strings"
)

type RequestUtils struct{}

type queryParameters struct {
	SVGType       *string
	GradientStart *string
	GradientEnd   *string
	TextColor     *string
	Text          *string
}

func NewRequestUtils() *RequestUtils {
	return &RequestUtils{}
}

func (ru *RequestUtils) GetIPAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	} else {
		ip = strings.Split(ip, ",")[0]
	}
	return ip
}

func (ru *RequestUtils) GetPageOwnerGitUsername(r *http.Request) string {
	githubProfileURL, err := url.Parse(r.Referer())
	if err != nil {
		return ""
	}

	pathSegments := strings.Split(githubProfileURL.Path, "/")
	if len(pathSegments) != 2 {
		return ""
	}
	return pathSegments[1]
}

func (ru *RequestUtils) ParseQueryParameters(r *http.Request) *queryParameters {
	query := r.URL.Query()

	return &queryParameters{
		GradientStart: getQueryValue(query, "gradientStart"),
		GradientEnd:   getQueryValue(query, "gradientEnd"),
		TextColor:     getQueryValue(query, "textColor"),
		Text:          getQueryValue(query, "text"),
	}
}

func getQueryValue(query url.Values, key string) *string {
	if value, exists := query[key]; exists && len(value) > 0 {
		return &value[0]
	}
	return nil
}
