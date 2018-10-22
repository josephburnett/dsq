package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
	"time"

	"github.com/josephburnett/dsq/pkg/html"
	"github.com/josephburnett/dsq/pkg/server"
	"github.com/josephburnett/dsq/pkg/types"
)

var (
	enableParallelSearch  bool
	parallelSearchBackend string
	client                *server.Client
)

func init() {
	flag.BoolVar(&enableParallelSearch, "enableParallelSearch", false, "turns on parallel search via rpc to backend")
	flag.StringVar(&parallelSearchBackend, "parallelSearchBackend", "localhost:8080", "address of the parallel search rpc server")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[len("/"):] != "" {
		http.Error(w, "file not found", http.StatusBadRequest)
		return
	}
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
		msg = append(msg, "Welcome to Dou Shou Qi!")
		msg = append(msg, "How to play: <a href=\"https://en.wikipedia.org/wiki/Jungle_(board_game)\">https://en.wikipedia.org/wiki/Jungle_(board_game)</a>")
		msg = append(msg, "Source code: <a href=\"https://github.com/josephburnett/dsq\">https://github.com/josephburnett/dsq</a>")
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
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("requested move %v on board\n%v\n", move, board)
		if winner := board.Winner(); winner != types.None {
			msg = append(msg, "Game over!")
			if winner == types.A {
				msg = append(msg, "Computer wins.")
			} else {
				msg = append(msg, "Human wins.")
			}
		} else {
			start := time.Now()
			reply, err := client.Move(board, move)
			if err != nil {
				http.Error(w, err.Error(), http.StatusServiceUnavailable)
				return
			}
			if winner := board.Winner(); winner != types.None {
				msg = append(msg, "Game over!")
				if winner == types.A {
					msg = append(msg, "Computer wins.")
				} else {
					msg = append(msg, "Human wins.")
				}
			} else {
				if reply.Ok {
					stat := reply.Stat
					msg = append(msg, fmt.Sprintf("Human moved %v.", move))
					msg = append(msg, fmt.Sprintf("Computer counter-moved %v.", reply.BestMove))
					msg = append(msg, fmt.Sprintf("Evaluated %v positions.", stat.PositionsEvaluated))
					msg = append(msg, fmt.Sprintf("Spent %v searching.", stat.Time))
					latency := time.Since(start)
					msg = append(msg, fmt.Sprintf("Request Latency %v.", latency))
					log.Printf("latency %v", latency)
					msg = append(msg, fmt.Sprintf("Best outcome is %v (positive is good).", -stat.BestOutcome))
					log.Printf("valid move returning stat %+v", stat)
				} else {
					msg = append(msg, fmt.Sprintf("Invalid move %v.", move))
					log.Printf("invalid move")
				}
			}
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

func imageHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/images/"):]
	file, err := html.Image(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(file)
}

func main() {
	flag.Parse()

	client = &server.Client{
		EnableParallelSearch:  enableParallelSearch,
		ParallelSearchBackend: parallelSearchBackend,
	}

	s := new(server.Search)
	rpc.Register(s)
	rpc.HandleHTTP()
	http.HandleFunc("/images/", imageHandler)
	http.HandleFunc("/", rootHandler)

	log.Printf("dsq up")
	log.Printf("enableParallelSearch=%v", enableParallelSearch)
	if enableParallelSearch {
		log.Printf("parallelSearchBackend=%v", parallelSearchBackend)
	}
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
