package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	port = ":8080"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	value := os.Getenv("WHO")
	if value == "" {
		value = "World"
	}
	fmt.Fprintf(w, "Hello %s\n", value)
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(port, nil)
}
