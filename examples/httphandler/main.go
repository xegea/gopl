package main

import (
	"fmt"
	"net/http"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello World!")
// }

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type helloWorldHandler struct{}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
}

func myHelloWorldFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloWorld Function")
}

func main() {

	// http.HandleFunc("/", handlerFunc)

	// var helloWorldHandler helloWorldHandler
	// http.Handle("/helloworld", helloWorldHandler)

	myHelloWorldFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HelloWorld Function")
	}
	http.HandleFunc("/helloworld", myHelloWorldFunc)

	fmt.Println("Listen on port :3000 ...")
	http.ListenAndServe(":3000", nil)
}
