package model

import (
	"fmt"
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"
	fmt.Fprintf(w, message)
}

func main2() {
	http.HandleFunc("/appHelloWorld", sayHelloWorld)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
