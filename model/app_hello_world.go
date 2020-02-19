package model

import (
	"fmt"
	"net/http"
	"strconv"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"
	fmt.Fprintf(w, message)
}

func launchHelloWorld(instance instance) {
	var port = strconv.Itoa(instance.port)
	mux := http.NewServeMux()
	mux.HandleFunc(instance.address, sayHelloWorld)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		panic(err)
	}
}