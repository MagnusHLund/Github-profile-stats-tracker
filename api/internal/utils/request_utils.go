package utils

import (
	"net/http"
	"strings"
)

type RequestUtils struct{}

func NewRequestUtils() *RequestUtils {
	return &RequestUtils{}
}

func GetIPAddress(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	} else {
		ip = strings.Split(ip, ",")[0]
	}
	return ip
}
