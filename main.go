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

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	agent := r.Header.Get("User-Agent")
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, agent)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.PostForm)
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
	http.HandleFunc("/headers", log(headers))
	http.HandleFunc("/body", log(body))
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
