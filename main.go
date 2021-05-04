package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", addHeaders(http.FileServer(http.Dir("video"))))

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil)
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Range")
		w.Header().Set("Content-Type", "video/mp4")
		w.Header().Set("mimetype", "application/dash+xml")
		h.ServeHTTP(w, r)
	}
}
