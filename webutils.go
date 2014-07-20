package main

import (
	"github.com/PuerkitoBio/ghost/handlers"
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
