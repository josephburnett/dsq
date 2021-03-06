package types

import (
	"testing"
)

var (
	down  = Point{0, -1}
	left  = Point{-1, 0}
	right = Point{1, 0}
	up    = Point{0, 1}
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
		point: Point{6, 0},
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
		move:          [2]Point{{6, 0}, {6, 1}},
		wantDisplaced: Empty,
		wantPlaced:    ATiger,
	}, {
		name:          "displace tiger with tiger",
		board:         NewBoard(),
		move:          [2]Point{{6, 0}, {0, 8}},
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
	}, {
		name:  "cat can take opposing dog on trap",
		board: EmptyBoard().With(Point{1, 0}, ACat).With(Point{2, 0}, BDog),
		want: [][2]Point{
			{{1, 0}, {0, 0}}, // cat left
			{{1, 0}, {2, 0}}, // cat right (takes dog)
			{{1, 0}, {1, 1}}, // cat up
			{{2, 0}, {1, 0}}, // dog left (takes cat)
			{{2, 0}, {3, 0}}, // dog right
			{{2, 0}, {2, 1}}, // dog up
		},
	}, {
		name:  "cat can not take own dog on trap",
		board: EmptyBoard().With(Point{1, 0}, ACat).With(Point{2, 0}, ADog),
		want: [][2]Point{
			{{1, 0}, {0, 0}}, // cat left
			{{1, 0}, {1, 1}}, // cat up
			{{2, 0}, {2, 1}}, // dog up
		},
	}, {
		name:  "can not move into own den",
		board: EmptyBoard().With(Point{2, 0}, ACat),
		want: [][2]Point{
			{{2, 0}, {1, 0}}, // left
			{{2, 0}, {2, 1}}, // up
		},
	}, {
		name: "lion cannot jump over mouse",
		board: EmptyBoard().With(
			Point{1, 3}, AMouse).With(
			Point{0, 3}, ALion),
		want: [][2]Point{
			{{0, 3}, {0, 2}}, // lion up
			{{0, 3}, {0, 4}}, // lion down
			{{1, 3}, {1, 2}}, // mouse up
			{{1, 3}, {2, 3}}, // mouse right
			{{1, 3}, {1, 4}}, // mouse down
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

func TestWinner(t *testing.T) {
	cases := []struct {
		name  string
		board *Board
		want  Side
	}{{
		name:  "new game not over",
		board: NewBoard(),
		want:  None,
	}, {
		name:  "a wins",
		board: NewBoard().With(BDen, ACat),
		want:  A,
	}, {
		name:  "b wins",
		board: NewBoard().With(ADen, BCat),
		want:  B,
	}, {
		name: "a wins when b has no moves",
		board: EmptyBoard().With(
			Point{0, 0}, BCat).With( // trapped in corner with no moves
			Point{1, 0}, ALion).With(
			Point{0, 1}, ATiger),
		want: A,
	}, {
		name: "b wins when a has no moves",
		board: EmptyBoard().With(
			Point{0, 0}, ACat).With( // trapped in corner with no moves
			Point{1, 0}, BLion).With(
			Point{0, 1}, BTiger),
		want: B,
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.board.Winner(); got != tc.want {
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
			want := 0
			for _, delta := range []Point{
				{0, -1}, // Down
				{-1, 0}, // Left
				{1, 0},  // Right
				{0, 1},  // Up
			} {
				to := Point{x + delta[0], y + delta[1]}
				if !isWater(from) && !isWater(to) && !isOutOfBounds(to) {
					want++
				}
				switch {
				case isWater(from) && len(adjacencies) != 0:
					t.Errorf("invalid normal adjacency from water: %v %v", from, to)
				case isOutOfBounds(to) && contains(adjacencies, to):
					t.Errorf("invalid normal adjacency with out-of-bounds: %v %v", from, to)
				case isWater(to) && contains(adjacencies, to):
					t.Errorf("invalid normal adjacency to water: %v %v", from, to)
				case !isOutOfBounds(to) && !isWater(from) && !isWater(to) && !contains(adjacencies, to):
					t.Errorf("missing normal adjacency: %v %v", from, to)
				}
			}
			if got := len(adjacencies); got != want {
				t.Errorf("Got %v normal adjacencies. Wanted %v. %v is adjacent to %v", got, want, from, adjacencies)
			}
		}
	}
}

func jump(from Point, delta Point) Point {
	switch from[0] {
	case 1, 2, 4, 5:
		if from[1] == 2 && delta == up {
			// Jumping up
			return Point{from[0], 6}
		}
		if from[1] == 6 && delta == down {
			// Jumping down
			return Point{from[0], 2}
		}
	}
	switch from[1] {
	case 3, 4, 5:
		if from[0] == 0 && delta == right {
			// Jumping right from left side
			return Point{3, from[1]}
		}
		if from[0] == 3 && delta == left {
			// Jumping left from center
			return Point{0, from[1]}
		}
		if from[0] == 3 && delta == right {
			// Jumping right from center
			return Point{6, from[1]}
		}
		if from[0] == 6 && delta == left {
			// Jumping left from right side
			return Point{3, from[1]}
		}
	}
	return Point{from[0] + delta[0], from[1] + delta[1]}
}

func TestJumpingAdjacency(t *testing.T) {
	for x := 0; x < 7; x++ {
		for y := 0; y < 9; y++ {
			from := Point{x, y}
			adjacencies := jumpingAdjacency[from]
			want := 0
			for _, delta := range []Point{
				{0, -1}, // Down
				{-1, 0}, // Left
				{1, 0},  // Right
				{0, 1},  // Up
			} {
				to := jump(from, delta)
				if !isWater(from) && !isWater(to) && !isOutOfBounds(to) {
					want++
				}
				switch {
				case isWater(from) && len(adjacencies) != 0:
					t.Errorf("invalid jumping adjacency from water: %v %v", from, to)
				case isOutOfBounds(to) && contains(adjacencies, to):
					t.Errorf("invalid jumping adjacency with out-of-bounds: %v %v", from, to)
				case isWater(to) && contains(adjacencies, to):
					t.Errorf("invalid jumping adjacency to water: %v %v", from, to)
				case !isOutOfBounds(to) && !isWater(from) && !isWater(to) && !contains(adjacencies, to):
					t.Errorf("missing jumping adjacency: %v %v", from, to)
				}
			}
			if got := len(adjacencies); got != want {
				t.Errorf("Got %v jumping adjacencies. Wanted %v. %v is adjacent to %v", got, want, from, adjacencies)
			}
		}
	}
}

func TestSwimmingAdjacency(t *testing.T) {
	for x := 0; x < 7; x++ {
		for y := 0; y < 9; y++ {
			from := Point{x, y}
			adjacencies := swimmingAdjacency[from]
			want := 0
			for _, delta := range []Point{
				{0, -1}, // Down
				{-1, 0}, // Left
				{1, 0},  // Right
				{0, 1},  // Up
			} {
				to := Point{x + delta[0], y + delta[1]}
				if !isOutOfBounds(to) {
					want++
				}
				switch {
				case isOutOfBounds(to) && contains(adjacencies, to):
					t.Errorf("invalid swimming adjacency with out-of-bounds: %v %v", from, to)
				case !isOutOfBounds(to) && !isWater(from) && !isWater(to) && !contains(adjacencies, to):
					t.Errorf("missing swimming adjacency: %v %v", from, to)
				}
			}
			if got := len(adjacencies); got != want {
				t.Errorf("Got %v swimming adjacencies. Wanted %v. %v is adjacent to %v", got, want, from, adjacencies)
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
