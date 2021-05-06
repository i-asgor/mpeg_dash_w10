package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/", addHeaders(http.FileServer(http.Dir("video"))))
	http.HandleFunc("/home", home)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("static"))))

	fmt.Println("Starting server on :8080")

	http.ListenAndServe(":8080", nil)
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Methods", "GET, OPTION, HEAD, PATCH, PUT, OST, DELETE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Range")
		// w.Header().Set("Content-Type", "video/mp4")
		// w.Header().Set("mimetype", "application/dash+xml")
		h.ServeHTTP(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}
