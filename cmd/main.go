package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "togglemaster-avaluation running")
}

func main() {

	http.HandleFunc("/health", health)

	fmt.Println("Server running on :8083")

	http.ListenAndServe(":8083", nil)
}
