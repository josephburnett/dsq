package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	msg := make([]string, 0)
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
				start := time.Now()
				cm, stat, ok := engine.BestMove(b.Clone(), types.A, 4)
				if ok {
					b.Move(cm)
				}
				msg = append(msg, fmt.Sprintf("Moved %v.", move))
				msg = append(msg, fmt.Sprintf("Counter-moved %v%v%v%v.", cm[0][0], cm[0][1], cm[1][0], cm[1][1]))
				msg = append(msg, fmt.Sprintf("Evaluated %v positions.", stat.PositionsEvaluated))
				msg = append(msg, fmt.Sprintf("Spent %v searching.", stat.Time))
				msg = append(msg, fmt.Sprintf("Latency %v.", time.Since(start)))
				msg = append(msg, fmt.Sprintf("Best outcome is %v.", -stat.BestOutcome))
				validMove = true
				break
			}
		}
		if !validMove {
			msg = append(msg, "Invalid move: "+move)
		}
	} else {
		msg = append(msg, "Game on!")
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
