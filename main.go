package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"
)

type Data struct {
	Info    string
	Results []float32
}

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

func process_file(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello World</h1></body>
	</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "Not implemented yet")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.Header().Set("dummy", "1")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Data{
		Info:    "done",
		Results: []float32{0.33, 0.33, 0.33},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "jkloe",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "asdf",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
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

	http.HandleFunc("/", root)
	http.HandleFunc("/hello", log(hello))
	http.HandleFunc("/world", log(world))
	http.HandleFunc("/headers", log(headers))
	http.HandleFunc("/body", log(body))
	http.HandleFunc("/process", log(process))
	http.HandleFunc("/process/file", log(process_file))
	http.HandleFunc("/write", log(writeExample))
	http.HandleFunc("/write/header", log(writeHeaderExample))
	http.HandleFunc("/redirect", log(headerExample))
	http.HandleFunc("/json", log(jsonExample))
	http.HandleFunc("/set_cookie", log(setCookie))
	http.HandleFunc("/get_cookie", log(getCookie))

	server.ListenAndServe()
}
