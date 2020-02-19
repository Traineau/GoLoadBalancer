package main

import (
	"fmt"
	"net/http"
)

func sayHello1(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/app1", sayHello1)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
