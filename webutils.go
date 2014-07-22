package main

import (
	"github.com/PuerkitoBio/ghost/handlers"
	"net"
	"net/http"
)

func PlainJSON(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func GZIPHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(handlers.GZIPHandler(h, nil))
}

func ParseIP(s string) string {
	ip, _, err := net.SplitHostPort(s)
	if err == nil {
		return ip
	}

	ip2 := net.ParseIP(s)
	if ip2 == nil {
		return ""
	}

	return ip2.String()
}

func CustomKeyGenerator(r *http.Request) string {
	return ParseIP(r.RemoteAddr)
}
