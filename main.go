package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	fmt.Println("got request")
}

func main() {
	http.HandleFunc("/", handler)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
