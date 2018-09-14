package main

import (
	"fmt"
	"net/http"

	"github.com/josephburnett/dsq-golang/pkg/html"
	"github.com/josephburnett/dsq-golang/pkg/types"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	move := r.Form.Get("move")
	if move != "" {
		fmt.Printf("move %v\n", move)
	}
	b := types.NewBoard()
	err = html.Render(w, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
