package main

import (
	"fmt"
	"net/http"
	"Backend/pkg/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.Hello)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}