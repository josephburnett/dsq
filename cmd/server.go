package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/josephburnett/dsq-golang/pkg/engine"
	"github.com/josephburnett/dsq-golang/pkg/html"
	"github.com/josephburnett/dsq-golang/pkg/types"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	move := r.Form.Get("move")
	board := r.Form.Get("board")
	b := types.NewBoard()
	msg := "Game on!"
	if move != "" && board != "" {
		requestedMove, err := parseMove(move)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		b, err = types.Unmarshal(board)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		validMove := false
		for _, m := range b.MoveList() {
			if m == requestedMove && b.Get(m[0]).Side() == types.B {
				b.Move(m)
				if counterMove, ok := engine.BestMove(b.Clone(), types.A); ok {
					b.Move(counterMove)
				}
				validMove = true
				break
			}
		}
		if validMove {
			msg = "Moved: " + move
		} else {
			msg = "Invalid move: " + move
		}
	}
	err = html.Render(w, b, msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
}

func parseMove(move string) ([2]types.Point, error) {
	coord := strings.Split(move, "")
	if len(coord) != 4 {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	fromX, err := strconv.Atoi(coord[0])
	if err != nil {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	fromY, err := strconv.Atoi(coord[1])
	if err != nil {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	toX, err := strconv.Atoi(coord[2])
	if err != nil {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	toY, err := strconv.Atoi(coord[3])
	if err != nil {
		return [2]types.Point{}, fmt.Errorf("Invalid move: %v", move)
	}
	return [2]types.Point{
		types.Point{fromX, fromY},
		types.Point{toX, toY},
	}, nil
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
