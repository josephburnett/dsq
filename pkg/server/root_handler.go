package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/josephburnett/dsq-golang/pkg/engine"
	"github.com/josephburnett/dsq-golang/pkg/html"
	"github.com/josephburnett/dsq-golang/pkg/types"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
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

				// Local:
				// cm, stat, ok := engine.BestMove(b.Clone(), types.A, 4)

				// Remote:
				cm, stat, ok, err := parallelBestMove(r.Host, b.Clone(), types.A, 5)
				if err != nil {
					http.Error(w, err.Error(), http.StatusServiceUnavailable)
					return
				}

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

func parallelBestMove(host string, b *types.Board, s types.Side, depth int) ([2]types.Point, *engine.Stat, bool, error) {
	reqCount := 0
	resChan := make(chan *searchResponse, 1000)
	stat := &engine.Stat{}
	haveMove := false
	var bestMove [2]types.Point
	var bestFitnessValue int
	for _, m := range b.MoveList() {
		side := b.Get(m[0]).Side()
		if s != side {
			continue
		}
		displaced := b.Move(m)
		other := types.A
		if s == types.A {
			other = types.B
		}
		req := &searchRequest{
			Board: b,
			Side:  other,
			Depth: depth - 1,
		}
		blob := req.marshal()
		go func(b []byte) {
			raw, err := http.Post("http://"+host+"/search", "", bytes.NewBuffer(blob))
			if err != nil {
				resChan <- &searchResponse{Err: err}
				return
			}
			bt, err := ioutil.ReadAll(raw.Body)
			if len(bt) == 0 {
				resChan <- &searchResponse{Err: fmt.Errorf("Empty response.")}
				return
			}
			if err != nil {
				resChan <- &searchResponse{Err: err}
				return
			}
			res, err := unmarshalResponse(bt)
			if err != nil {
				resChan <- &searchResponse{Err: err}
				return
			}
			resChan <- res
		}(blob)
		b.Unmove(m, displaced)
		reqCount++
	}
	var err error
	for i := 0; i < reqCount; i++ {
		res := <-resChan
		if res.Err != nil {
			err = res.Err
			continue
		}
		if s == types.A {
			min := res.Stat.BestOutcome
			// Choose the maximum of the minimums
			if min > bestFitnessValue || !haveMove {
				bestMove = res.BestMove
				bestFitnessValue = min
				haveMove = true
			}
		} else {
			max := res.Stat.BestOutcome
			// Choose the minimum of the maximums
			if max < bestFitnessValue || !haveMove {
				bestMove = res.BestMove
				bestFitnessValue = max
				haveMove = true
			}
		}
		stat.Time = stat.Time + res.Stat.Time
		stat.PositionsEvaluated = stat.PositionsEvaluated + res.Stat.PositionsEvaluated
	}
	return bestMove, stat, haveMove, err
}
