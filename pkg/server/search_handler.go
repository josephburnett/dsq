package server

import (
	"io/ioutil"
	"net/http"

	"github.com/josephburnett/dsq-golang/pkg/engine"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req, err := unmarshalRequest(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	bestMove, stat, ok := engine.BestMove(req.Board, req.Side, req.Depth)
	res := &searchResponse{
		BestMove: bestMove,
		Stat:     stat,
		Ok:       ok,
	}
	w.Write(res.marshal())
}
