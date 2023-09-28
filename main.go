package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!!!\n")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	log.Fatal(http.ListenAndServe(":1337", nil))
}
