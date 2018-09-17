package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
	"time"

	"github.com/josephburnett/dsq-golang/pkg/html"
	"github.com/josephburnett/dsq-golang/pkg/server"
	"github.com/josephburnett/dsq-golang/pkg/types"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	msg := make([]string, 0)
	switch r.Method {
	case http.MethodGet:
		msg = append(msg, "Game on!")
		err = html.Render(w, types.NewBoard(), msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
	case http.MethodPost:
		move, err := parseMove(r.Form.Get("move"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		board, err := types.Unmarshal(r.Form.Get("board"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		start := time.Now()
		reply, err := server.Move(r.Host, board, move)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		if reply.Ok {
			stat := reply.Stat
			msg = append(msg, fmt.Sprintf("Moved %v.", move))
			msg = append(msg, fmt.Sprintf("Counter-moved %v.", reply.BestMove))
			msg = append(msg, fmt.Sprintf("Evaluated %v positions.", stat.PositionsEvaluated))
			msg = append(msg, fmt.Sprintf("Spent %v searching.", stat.Time))
			msg = append(msg, fmt.Sprintf("Latency %v.", time.Since(start)))
			msg = append(msg, fmt.Sprintf("Best outcome is %v.", -stat.BestOutcome))
		} else {
			msg = append(msg, fmt.Sprintf("Invalid move %v.", move))
		}
		err = html.Render(w, board, msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	s := new(server.Search)
	rpc.Register(s)
	rpc.HandleHTTP()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func parseMove(move string) ([2]types.Point, error) {
	if move == "" {
		return [2]types.Point{}, fmt.Errorf("Empty move: %v", move)
	}
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
