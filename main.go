package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root url")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", log(root))
	http.HandleFunc("/hello", log(hello))
	http.HandleFunc("/world", log(world))

	server.ListenAndServe()
}
