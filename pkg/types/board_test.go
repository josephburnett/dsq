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
		board: EmptyBoard().With(Point{0, 0}, ATiger),
		want: [][2]Point{
			{{0, 0}, {1, 0}}, // right
			{{0, 0}, {0, 1}}, // up
		},
	}, {
		name:  "tiger by the water (bottom)",
		board: EmptyBoard().With(Point{1, 2}, BTiger),
		want: [][2]Point{
			{{1, 2}, {1, 1}}, // down
			{{1, 2}, {0, 2}}, // left
			{{1, 2}, {2, 2}}, // right
			{{1, 2}, {1, 6}}, // jump up over water
		},
	}, {
		name:  "tiger by the water (side)",
		board: EmptyBoard().With(Point{0, 4}, ATiger),
		want: [][2]Point{
			{{0, 4}, {0, 3}}, // down
			{{0, 4}, {3, 4}}, // jump right over water
			{{0, 4}, {0, 5}}, // up
		},
	}, {
		name:  "mouse by the water",
		board: EmptyBoard().With(Point{1, 2}, BMouse),
		want: [][2]Point{
			{{1, 2}, {1, 1}}, // down
			{{1, 2}, {0, 2}}, // left
			{{1, 2}, {2, 2}}, // right
			{{1, 2}, {1, 3}}, // up into the water
		},
	}, {
		name:  "mouse in the water",
		board: EmptyBoard().With(Point{1, 3}, AMouse),
		want: [][2]Point{
			{{1, 3}, {1, 2}}, // down out of the water
			{{1, 3}, {0, 3}}, // left out of the water
			{{1, 3}, {2, 3}}, // right
			{{1, 3}, {1, 4}}, // up
		},
	}, {
		name:  "cat in the middle",
		board: EmptyBoard().With(Point{3, 4}, BCat),
		want: [][2]Point{
			{{3, 4}, {3, 3}}, // down
			{{3, 4}, {3, 5}}, // up
		},
	}, {
		name:  "cat takes a mouse",
		board: EmptyBoard().With(Point{0, 0}, ACat).With(Point{0, 1}, BMouse),
		want: [][2]Point{
			{{0, 0}, {1, 0}}, // cat moves right
			{{0, 0}, {0, 1}}, // cat moves up taking mouse
			{{0, 1}, {1, 1}}, // mouse moves right
			{{0, 1}, {0, 2}}, // mouse moves up
		},
	}, {
		name:  "cat does not take a mouse",
		board: EmptyBoard().With(Point{0, 0}, ACat).With(Point{0, 1}, AMouse),
		want: [][2]Point{
			{{0, 0}, {1, 0}}, // cat moves right
			{{0, 1}, {1, 1}}, // mouse moves right
			{{0, 1}, {0, 2}}, // mouse moves up
		},
	}, {
		name:  "mouse on the opposing side",
		board: EmptyBoard().With(Point{3, 7}, AMouse),
		want: [][2]Point{
			{{3, 7}, {3, 6}}, // down
			{{3, 7}, {2, 7}}, // left
			{{3, 7}, {4, 7}}, // right
			{{3, 7}, {3, 8}}, // up
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

func contains(points []Point, p1 Point) bool {
	for _, p2 := range points {
		if p1 == p2 {
			return true
		}
	}
	return false
}

func isWater(p Point) bool {
	switch p[0] {
	case 1, 2, 4, 5:
		switch p[1] {
		case 3, 4, 5:
			return true
		default:
			return false
		}
	default:
		return false
	}
}

func isOutOfBounds(p Point) bool {
	if p[0] < 0 || p[0] > 6 {
		return true
	}
	if p[1] < 0 || p[1] > 8 {
		return true
	}
	return false
}

func TestNormalAdjacency(t *testing.T) {
	for x := 0; x < 7; x++ {
		for y := 0; y < 9; y++ {
			from := Point{x, y}
			adjacencies := normalAdjacency[from]
			for _, delta := range []Point{
				{0, -1}, // Down
				{-1, 0}, // Left
				{1, 0},  // Right
				{0, 1},  // Up
			} {
				to := Point{x + delta[0], y + delta[1]}
				switch {
				case isOutOfBounds(to) && contains(adjacencies, to):
					t.Errorf("invalid normal adjacency with out-of-bounds: %v %v", from, to)
				case isWater(to) && contains(adjacencies, to):
					t.Errorf("invalid normal adjacency with water: %v %v", from, to)
				case !isOutOfBounds(to) && !isWater(from) && !isWater(to) && !contains(adjacencies, to):
					t.Errorf("missing normal adjacency: %v %v", from, to)
				}
			}
		}
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
