package types

import (
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		name  string
		board *Board
		point Point
		want  Piece
	}{{
		name:  "get empty",
		board: NewBoard(),
		point: Point{0, 1},
		want:  Empty,
	}, {
		name:  "get tiger",
		board: NewBoard(),
		point: Point{0, 0},
		want:  ATiger,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.board.Get(tc.point); got != tc.want {
				t.Errorf("%q expected %v. got %v.", tc.name, tc.want, got)
			}
		})
	}
}

func TestMove(t *testing.T) {
	cases := []struct {
		name          string
		board         *Board
		move          [2]Point
		wantDisplaced Piece
		wantPlaced    Piece
	}{{
		name:          "move tiger to empty",
		board:         NewBoard(),
		move:          [2]Point{{0, 0}, {0, 1}},
		wantDisplaced: Empty,
		wantPlaced:    ATiger,
	}, {
		name:          "displace tiger with tiger",
		board:         NewBoard(),
		move:          [2]Point{{0, 0}, {6, 8}},
		wantDisplaced: BTiger,
		wantPlaced:    ATiger,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.board.Move(tc.move); got != tc.wantDisplaced {
				t.Errorf("%q expected to displace %v. got %v.", tc.name, tc.wantDisplaced, got)
			}
			if got := tc.board.Get(tc.move[1]); got != tc.wantPlaced {
				t.Errorf("%q expected to place %v. got %v.", tc.name, tc.wantPlaced, got)
			}
			if got := tc.board.Get(tc.move[0]); got != Empty {
				t.Errorf("%q expected an Empty space behind. got %v.", tc.name, got)
			}
		})
	}
}

func TestMoveList(t *testing.T) {
	cases := []struct {
		name  string
		board *Board
		want  [][2]Point
	}{{
		name:  "tiger in corner",
		board: emptyBoard().with(Point{0, 0}, ATiger),
		want: [][2]Point{
			{{0, 0}, {1, 0}}, // right
			{{0, 0}, {0, 1}}, // up
		},
	}, {
		name:  "tiger by the water (bottom)",
		board: emptyBoard().with(Point{1, 2}, BTiger),
		want: [][2]Point{
			{{1, 2}, {1, 1}}, // down
			{{1, 2}, {0, 2}}, // left
			{{1, 2}, {2, 2}}, // right
			{{1, 2}, {1, 6}}, // jump up over water
		},
	}, {
		name:  "tiger by the water (side)",
		board: emptyBoard().with(Point{0, 4}, ATiger),
		want: [][2]Point{
			{{0, 4}, {0, 3}}, // down
			{{0, 4}, {3, 4}}, // jump right over water
			{{0, 4}, {0, 5}}, // up
		},
	}, {
		name:  "mouse by the water",
		board: emptyBoard().with(Point{1, 2}, BMouse),
		want: [][2]Point{
			{{1, 2}, {1, 1}}, // down
			{{1, 2}, {0, 2}}, // left
			{{1, 2}, {2, 2}}, // right
			{{1, 2}, {1, 3}}, // up into the water
		},
	}, {
		name:  "mouse in the water",
		board: emptyBoard().with(Point{1, 3}, AMouse),
		want: [][2]Point{
			{{1, 3}, {1, 2}}, // down out of the water
			{{1, 3}, {0, 3}}, // left out of the water
			{{1, 3}, {2, 3}}, // right
			{{1, 3}, {1, 4}}, // up
		},
	}, {
		name:  "cat in the middle",
		board: emptyBoard().with(Point{3, 4}, BCat),
		want: [][2]Point{
			{{3, 4}, {3, 3}}, // down
			{{3, 4}, {3, 5}}, // up
		},
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.board.MoveList(); !equals(got, tc.want) {
				t.Errorf("%q expected %v. got %v.", tc.name, tc.want, got)
			}
		})
	}
}

func equals(a, b [][2]Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func emptyBoard() *Board {
	return &Board{}
}

func (b *Board) with(pt Point, p Piece) *Board {
	b.put(pt, p)
	return b
}
