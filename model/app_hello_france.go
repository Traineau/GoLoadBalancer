package model

import (
	"fmt"
	"net/http"
)

func sayHelloFrance(w http.ResponseWriter, r *http.Request) {
	message := "Hello France"
	fmt.Fprintf(w, message)
}

func launchHelloFrance(instance instance) {
	http.HandleFunc(instance.address, sayHelloFrance)
	if err := http.ListenAndServe(instance.port, nil); err != nil {
		panic(err)
	}
}

