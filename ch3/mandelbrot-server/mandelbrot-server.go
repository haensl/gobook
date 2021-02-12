package main

import (
	"haensl/gobook/ch3/mandelbrot"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		mandelbrot.Mandelbrot(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
