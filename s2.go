package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World from s2!")
	})

	fmt.Println("Server is running on http://localhost:8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
