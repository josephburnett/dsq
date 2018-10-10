package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/josephburnett/dsq/pkg/html"
	"github.com/josephburnett/dsq/pkg/server"
	"github.com/josephburnett/dsq/pkg/types"
)

var (
	backend = os.Getenv("BACKEND")
)

func init() {
	if backend == "" {
		backend = "localhost:8080"
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("frontend request")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	msg := make([]string, 0)
	switch r.Method {
	case http.MethodGet:
		b := types.NewBoard()
		msg = append(msg, "Game on!")
		err = html.Render(w, b, msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		log.Printf("new board\n%v\n", b)
	case http.MethodPost:
		move, err := parseMove(r.Form.Get("move"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		board, err := types.Unmarshal(r.Form.Get("board"))
		log.Printf("requested move %v on board\n%v\n", move, board)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		start := time.Now()
		reply, err := server.Move(backend, board, move)
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
			latency := time.Since(start)
			msg = append(msg, fmt.Sprintf("Latency %v.", latency))
			log.Printf("latency %v", latency)
			msg = append(msg, fmt.Sprintf("Best outcome is %v.", -stat.BestOutcome))
			log.Printf("valid move returning stat %+v", stat)
		} else {
			msg = append(msg, fmt.Sprintf("Invalid move %v.", move))
			log.Printf("invalid move")
		}
		err = html.Render(w, board, msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		log.Printf("updated board\n%v\n", board)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		log.Printf("method not allowed %", r.Method)
		return
	}
}

func main() {
	s := new(server.Search)
	rpc.Register(s)
	rpc.HandleHTTP()
	http.HandleFunc("/", handler)

	log.Printf("dsq up with backend %v", backend)
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
