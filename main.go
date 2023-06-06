package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloMessage() string {
	return "Hello"
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, HelloMessage())
}

func main() {
	port := ":9091"
	http.HandleFunc("/hello", HandleHello)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
