package main

import (
	"net/http"

	"github.com/josephburnett/dsq-golang/pkg/server"
)

func main() {
	http.HandleFunc("/", server.RootHandler)
	http.ListenAndServe(":8080", nil)
}
