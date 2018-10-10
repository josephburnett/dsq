package engine

import (
	"fmt"
	"testing"

	"github.com/josephburnett/dsq/pkg/types"
)

func TestBestMove(t *testing.T) {
	cases := []struct {
		name  string
		board *types.Board
		side  types.Side
		depth int
		want  [2]types.Point
	}{{
		name:  "mouse moves into opposing den",
		board: types.EmptyBoard().With(types.Point{3, 7}, types.AMouse),
		side:  types.A,
		depth: 1,
		want:  [2]types.Point{{3, 7}, {3, 8}},
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got, _, _ := BestMove(tc.board, tc.side, tc.depth); got != tc.want {
				fmt.Printf("move list %v\n", tc.board.MoveList())
				t.Errorf("%q expected %v. got %v.", tc.name, tc.want, got)
			}
		})
	}
}

func TestNoSideEffects(t *testing.T) {
	got := types.NewBoard()
	want := types.NewBoard()
	BestMove(got, types.A, 2)
	if *got != *want {
		t.Errorf("\nexpected \n%v\ngot \n%v\n", want, got)
	}
}
