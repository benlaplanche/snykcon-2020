package main

import (
	"fmt"
	"net/http"
)

// HelloSnykcon exported for tests
func HelloSnykcon(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello Snykcon 2020!!!")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", HelloSnykcon)
	http.HandleFunc("/hello", HelloSnykcon)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}
