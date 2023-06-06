package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	port := ":9091"
	http.HandleFunc("/hello", handleHello)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
