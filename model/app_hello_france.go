package model

import (
	"fmt"
	"net/http"
	"strconv"
)

func sayHelloFrance(w http.ResponseWriter, r *http.Request) {
	message := "Hello France"
	fmt.Fprintf(w, message)
}

func launchHelloFrance(instance instance) {
	var port = strconv.Itoa(instance.port)
	mux := http.NewServeMux()
	mux.HandleFunc(instance.address, sayHelloFrance)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		panic(err)
	}
}

