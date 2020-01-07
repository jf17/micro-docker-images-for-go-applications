package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
  http.HandleFunc("/jf17", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, jf17.ru !")
	})
	http.ListenAndServe(":3000", nil)
}
