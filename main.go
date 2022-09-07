package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message string = "Hello, you've requested: " + r.URL.Path
		if r.URL.Path != "/favicon.ico" {
			fmt.Println(message)
		}
		fmt.Fprintf(w, message)
	})

	http.ListenAndServe(":8080", nil)
}
