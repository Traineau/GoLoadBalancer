package model

import (
	"fmt"
	"net/http"
)

func sayHelloFrance(w http.ResponseWriter, r *http.Request) {
	message := "Hello France"
	fmt.Fprintf(w, message)
}

func main1() {
	http.HandleFunc("/appHelloFrance", sayHelloFrance)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
