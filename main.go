package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloMessage() string {
	return "Hello"
}

func HandleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, HelloMessage())
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/hello", HandleHello)
	fmt.Println("aplicação ouvindo na porta ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
