package server

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/josephburnett/dsq-golang/pkg/engine"
	"github.com/josephburnett/dsq-golang/pkg/types"
)

type Search int

type Request struct {
	Board *types.Board
	Side  types.Side
	Depth int
}

type Reply struct {
	BestMove         [2]types.Point
	BestFitnessValue int
	Stat             *engine.Stat
	Ok               bool
}

func (s *Search) BestMove(req Request, res *Reply) error {
	log.Printf("search request")
	if req.Board == nil {
		return fmt.Errorf("board is required")
	}
	if req.Side != types.A && req.Side != types.B {
		return fmt.Errorf("invalid side %v", req.Side)
	}
	if req.Depth < 1 || req.Depth > 6 {
		return fmt.Errorf("invalid depth %v", req.Depth)
	}
	bestMove, stat, ok := engine.BestMove(req.Board, req.Side, req.Depth)
	res.BestMove = bestMove
	res.Stat = stat
	res.Ok = ok
	log.Printf("search successful")
	return nil
}

func ParallelSearch(address string, r Request) (*Reply, error) {
	log.Printf("starting parallel search")
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Printf("error connecting to backend %v", err)
		return nil, err
	}
	calls := make([]*rpc.Call, 0)
	i := 0
	for _, m := range r.Board.MoveList() {
		side := r.Board.Get(m[0]).Side().Other()
		if side != r.Side {
			continue
		}
		b := r.Board.Clone()
		b.Move(m)
		req := Request{
			Board: b,
			Side:  r.Side,
			Depth: r.Depth - 1,
		}
		reply := &Reply{}
		log.Printf("making search request %v", i)
		call := client.Go("Search.BestMove", req, reply, nil)
		calls = append(calls, call)
		i++
	}
	haveMove := false
	var bestMove [2]types.Point
	var bestFitnessValue int
	stat := &engine.Stat{}
	for i, c := range calls {
		log.Printf("waiting on search request %v", i)
		call := <-c.Done
		if call.Error != nil {
			log.Printf("search request error %v", err)
			return nil, err
		}
		reply := c.Reply.(*Reply)
		if r.Side == types.A {
			min := reply.BestFitnessValue
			// Choose the maximum of the minimums
			if min > bestFitnessValue || !haveMove {
				bestMove = reply.BestMove
				bestFitnessValue = reply.BestFitnessValue
				haveMove = true
			}
		} else {
			max := reply.BestFitnessValue
			// Choose the minimum of the maximums
			if max < bestFitnessValue || !haveMove {
				bestMove = reply.BestMove
				bestFitnessValue = reply.BestFitnessValue
				haveMove = true
			}
		}
		stat.Time = stat.Time + reply.Stat.Time
		stat.PositionsEvaluated = stat.PositionsEvaluated + reply.Stat.PositionsEvaluated
		log.Printf("search request %v successful", i)
	}
	return &Reply{
		BestMove:         bestMove,
		BestFitnessValue: bestFitnessValue,
		Stat:             stat,
		Ok:               haveMove,
	}, nil
}
