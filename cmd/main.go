package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "togglemaster-evaluation running")
}

func main() {

	http.HandleFunc("/health", health)

	fmt.Println("Server running on :8083")

	err := http.ListenAndServe(":8083", nil)

	if err != nil {
		panic(err)
	}
}
