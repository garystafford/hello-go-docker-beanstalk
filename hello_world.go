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
  fmt.Fprintf(w, "Hello, world! You have called me %d times.\n", calls)
}

func handlerICon(w http.ResponseWriter, r *http.Request) {}

func init() {
  fmt.Printf("Started server at http://localhost%v.\n", port)
  http.HandleFunc("/favicon.ico", handlerICon)
  http.HandleFunc("/", helloWorld)
  http.ListenAndServe(port, nil)
}

func main() {}
