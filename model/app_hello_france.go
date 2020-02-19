package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello France"
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/app2", sayHello)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
