package model

import (
	"fmt"
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"
	fmt.Fprintf(w, message)
}

func launchHelloWorld(instance instance) {
	http.HandleFunc(instance.address, sayHelloWorld)
	if err := http.ListenAndServe(instance.port, nil); err != nil {
		panic(err)
	}
}
