package engine

import (
	"testing"

	"github.com/josephburnett/dsq/pkg/types"
)

func TestFitness(t *testing.T) {
	cases := []struct {
		name  string
		board *types.Board
		want  int
	}{{
		name:  "single a-side mouse",
		board: types.EmptyBoard().With(types.Point{0, 0}, types.AMouse),
		want:  500 + 8,
	}, {
		name:  "single b-side mouse",
		board: types.EmptyBoard().With(types.Point{6, 8}, types.BMouse),
		want:  -(500 + 8),
	}, {
		name:  "winning a-side mouse",
		board: types.EmptyBoard().With(types.Point{3, 8}, types.AMouse),
		want:  500 + 9999,
	}, {
		name:  "winning b-side mouse",
		board: types.EmptyBoard().With(types.Point{3, 0}, types.BMouse),
		want:  -(500 + 9999),
	}, {
		name:  "neutral starting point",
		board: types.NewBoard(),
		want:  0,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Fitness(tc.board); got != tc.want {
				t.Errorf("%q expected %v. got %v.", tc.name, tc.want, got)
			}
		})
	}
}
