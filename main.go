package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Infof("request method: %s", req.Method)
	log.Infof("request url: %s", req.URL)
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	log.Infof("request method: %s", req.Method)
	log.Infof("request url: %s", req.URL)
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}
