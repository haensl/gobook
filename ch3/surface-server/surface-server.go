package main

import (
	"haensl/gobook/ch3/surface"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		surface.Surface(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
