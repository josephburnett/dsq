package server

import (
	"github.com/josephburnett/dsq/pkg/engine"
	"github.com/josephburnett/dsq/pkg/types"
)

type Client struct {
	EnableParallelSearch  bool
	ParallelSearchBackend string
}

func (c *Client) Move(b *types.Board, move [2]types.Point) (*Reply, error) {
	for _, m := range b.MoveList() {
		// Only accept valid moves by player B (human)
		if m == move && b.Get(m[0]).Side() == types.B {
			b.Move(m)
			if c.EnableParallelSearch {
				// Send parallel RPCs to search server
				reply, err := ParallelSearch(c.ParallelSearchBackend, Request{
					Board: b.Clone(),
					Side:  types.A,
					Depth: 5,
				})
				if err != nil {
					return nil, err
				}
				if reply.Ok {
					b.Move(reply.BestMove)
				}
				return reply, nil
			} else {
				// Use the engine directly
				bestMove, stat, ok := engine.BestMove(b.Clone(), types.A, 5)
				if ok {
					b.Move(bestMove)
					reply := &Reply{
						BestMove:         bestMove,
						BestFitnessValue: stat.BestOutcome,
						Stat:             stat,
						Ok:               ok,
					}
					return reply, nil
				}
			}
		}
	}
	return &Reply{}, nil
}
