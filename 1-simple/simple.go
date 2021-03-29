package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	mediaDir = "1-simple/media"
	port     = 8080
)

func main() {
	http.Handle("/", handle(http.FileServer(http.Dir(mediaDir))))
	log.Printf("Serving %s at http://localhost:%d/outputlist.m3u8\n", mediaDir, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func handle(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
