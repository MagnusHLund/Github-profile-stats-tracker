package utils

import (
	"net/http"
	"net/url"
	"strings"
)

type RequestUtils struct{}

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

func (ph *RequestUtils) GetPageOwnerGitUsername(r *http.Request) string {
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
