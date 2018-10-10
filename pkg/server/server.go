package server

import (
	"github.com/josephburnett/dsq/pkg/types"
)

func Move(host string, b *types.Board, move [2]types.Point) (*Reply, error) {
	for _, m := range b.MoveList() {
		if m == move && b.Get(m[0]).Side() == types.B {
			b.Move(m)
			reply, err := ParallelSearch(host, Request{
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
		}
	}
	return &Reply{}, nil
}
