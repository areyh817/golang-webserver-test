package main

import (
	"fmt"
	"net/http"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test !!!"))
	})

	http.HandleFunc("/bar", barHandler)

	http.ListenAndServe(":5000", nil)
}
