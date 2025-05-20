package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello Go!")
	http.HandleFunc("/", handleProxy)

	log.Println("Starting proxy server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
