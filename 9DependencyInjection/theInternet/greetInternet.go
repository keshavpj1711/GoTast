package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GreetInternet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	// here http.ResponseWriter acts as a io.Writer
	GreetInternet(w, "Internet")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
