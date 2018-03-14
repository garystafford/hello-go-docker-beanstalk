package main

import (
	"fmt"
	"net/http"
)

const (
	port = ":3000"
)

var calls = 0

func helloWorld(w http.ResponseWriter, r *http.Request) {
	calls++
	fmt.Fprintf(w, "Hello AWS Elastic Beanstalk: v2. You have called me %d times.\n", calls)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
}

func init() {
	fmt.Printf("Started server at http://localhost%v.\n", port)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(port, nil)
}

func main() {}
