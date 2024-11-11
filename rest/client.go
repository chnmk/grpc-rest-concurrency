package rest_api

import (
	"io"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/chnmk/grpc-rest-concurrency/counter"
)

func Client() {
	r, err := http.Get("http://localhost:8080/rest")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if string(body) == "Hello, World!" {
		atomic.AddInt32(&counter.Counter_REST, 1)
	} else {
		log.Fatal("Wrong response body: ", body)
	}
}
