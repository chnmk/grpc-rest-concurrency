package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func restHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func rest() {
	http.HandleFunc("/rest", restHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func grpc() {

}

func main() {
	go rest()
	go grpc()

	runtime.Goexit()
}
