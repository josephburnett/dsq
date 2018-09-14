package types

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Board [9][7]Piece
type Point [2]int

var (
	ADen = Point{3, 0}
	BDen = Point{3, 8}
)

func NewBoard() *Board {
	return &Board{
		// Row 0
		{ATiger, Empty, Empty, Empty, Empty, Empty, ALion},
		// Row 1
		{Empty, ACat, Empty, Empty, Empty, ADog, Empty},
		// Row 2
		{AElephant, Empty, AWolf, Empty, AHyena, Empty, AMouse},
		// Row 3
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		// Row 4
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		// Row 5
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		// Row 6
		{BMouse, Empty, BHyena, Empty, BCat, Empty, BElephant},
		// Row 7
		{Empty, BDog, Empty, Empty, Empty, BCat, Empty},
		// Row 8
		{BLion, Empty, Empty, Empty, Empty, Empty, BTiger},
	}
}

func (b *Board) Get(p Point) Piece {
	return b[p[1]][p[0]]
}

func (b *Board) put(p Point, piece Piece) Piece {
	displaced := b[p[1]][p[0]]
	b[p[1]][p[0]] = piece
	return displaced
}

func (b *Board) Move(m [2]Point) Piece {
	return b.put(m[1], b.put(m[0], Empty))
}

func (b *Board) Unmove(m [2]Point, p Piece) {
	b.put(m[1], b.put(m[0], p))
}

func (b *Board) MoveList() [][2]Point {
	moves := make([][2]Point, 0)
	for x := 0; x < 7; x++ {
		for y := 0; y < 9; y++ {
			from := Point{x, y}
			p := b.Get(from)
			if p == Empty {
				continue
			}
			var adjacency map[Point][]Point
			switch {
			case p.CanJump():
				adjacency = jumpingAdjacency
			case p.CanSwim():
				adjacency = swimmingAdjacency
			default:
				adjacency = normalAdjacency
			}
			for _, to := range adjacency[from] {
				opponent := b.Get(to)
				if p.CanTake(opponent) {
					moves = append(moves, [2]Point{from, to})
				}
			}
		}
	}
	return moves
}

func (p Point) Rotate() Point {
	return Point{6 - p[0], 8 - p[1]}
}

var normalAdjacency = map[Point][]Point{
	// Row 0
	Point{0, 0}: []Point{Point{1, 0}, Point{0, 1}},
	Point{1, 0}: []Point{Point{0, 0}, Point{2, 0}, Point{1, 1}},
	Point{2, 0}: []Point{Point{1, 0}, Point{3, 0}, Point{2, 1}},
	Point{3, 0}: []Point{Point{2, 0}, Point{4, 0}, Point{3, 1}},
	Point{4, 0}: []Point{Point{3, 0}, Point{5, 0}, Point{4, 1}},
	Point{5, 0}: []Point{Point{4, 0}, Point{6, 0}, Point{5, 1}},
	Point{6, 0}: []Point{Point{5, 0}, Point{6, 1}},
	// Row 1
	Point{0, 1}: []Point{Point{0, 0}, Point{1, 1}, Point{0, 2}},
	Point{1, 1}: []Point{Point{1, 0}, Point{0, 1}, Point{2, 1}, Point{1, 2}},
	Point{2, 1}: []Point{Point{2, 0}, Point{1, 1}, Point{3, 1}, Point{2, 2}},
	Point{3, 1}: []Point{Point{3, 0}, Point{2, 1}, Point{4, 1}, Point{3, 2}},
	Point{4, 1}: []Point{Point{4, 0}, Point{3, 1}, Point{5, 1}, Point{4, 2}},
	Point{5, 1}: []Point{Point{5, 0}, Point{4, 1}, Point{6, 1}, Point{5, 2}},
	Point{6, 1}: []Point{Point{6, 0}, Point{5, 1}, Point{6, 2}},
	// Row 2
	Point{0, 2}: []Point{Point{0, 1}, Point{1, 2}, Point{0, 3}},
	Point{1, 2}: []Point{Point{1, 1}, Point{0, 2}, Point{2, 2}},
	Point{2, 2}: []Point{Point{2, 1}, Point{1, 2}, Point{3, 2}},
	Point{3, 2}: []Point{Point{3, 1}, Point{2, 2}, Point{4, 2}, Point{3, 3}},
	Point{4, 2}: []Point{Point{4, 1}, Point{3, 2}, Point{5, 2}},
	Point{5, 2}: []Point{Point{5, 1}, Point{4, 2}, Point{6, 2}},
	Point{6, 2}: []Point{Point{6, 1}, Point{5, 2}, Point{6, 3}},
	// Row 3
	Point{0, 3}: []Point{Point{0, 2}, Point{0, 4}},
	Point{1, 3}: []Point{},
	Point{2, 3}: []Point{},
	Point{3, 3}: []Point{Point{3, 2}, Point{3, 4}},
	Point{4, 3}: []Point{},
	Point{5, 3}: []Point{},
	Point{6, 3}: []Point{Point{6, 2}, Point{6, 4}},
	// Row 4
	Point{0, 4}: []Point{Point{0, 3}, Point{0, 5}},
	Point{1, 4}: []Point{},
	Point{2, 4}: []Point{},
	Point{3, 4}: []Point{Point{3, 3}, Point{3, 5}},
	Point{4, 4}: []Point{},
	Point{5, 4}: []Point{},
	Point{6, 4}: []Point{Point{6, 3}, Point{6, 5}},
	// Row 5
	Point{0, 5}: []Point{Point{0, 4}, Point{0, 6}},
	Point{1, 5}: []Point{},
	Point{2, 5}: []Point{},
	Point{3, 5}: []Point{Point{3, 4}, Point{3, 6}},
	Point{4, 5}: []Point{},
	Point{5, 5}: []Point{},
	Point{6, 5}: []Point{Point{6, 4}, Point{6, 6}},
	// Row 6
	Point{0, 6}: []Point{Point{0, 5}, Point{1, 6}, Point{0, 7}},
	Point{1, 6}: []Point{Point{0, 6}, Point{2, 6}, Point{1, 7}},
	Point{2, 6}: []Point{Point{1, 6}, Point{3, 6}, Point{2, 7}},
	Point{3, 6}: []Point{Point{3, 5}, Point{2, 6}, Point{4, 6}, Point{3, 7}},
	Point{4, 6}: []Point{Point{3, 6}, Point{5, 6}, Point{4, 7}},
	Point{5, 6}: []Point{Point{4, 6}, Point{6, 6}, Point{5, 7}},
	Point{6, 6}: []Point{Point{6, 5}, Point{5, 6}, Point{6, 7}},
	// Row 7
	Point{0, 7}: []Point{Point{0, 6}, Point{1, 7}, Point{0, 8}},
	Point{1, 7}: []Point{Point{1, 6}, Point{0, 7}, Point{2, 7}, Point{1, 8}},
	Point{2, 7}: []Point{Point{2, 6}, Point{1, 7}, Point{3, 7}, Point{2, 8}},
	Point{3, 7}: []Point{Point{3, 6}, Point{2, 7}, Point{4, 7}, Point{3, 8}},
	Point{4, 7}: []Point{Point{4, 6}, Point{3, 7}, Point{5, 7}, Point{4, 8}},
	Point{5, 7}: []Point{Point{5, 6}, Point{4, 7}, Point{6, 7}, Point{5, 8}},
	Point{6, 7}: []Point{Point{6, 6}, Point{5, 7}, Point{6, 8}},
	// Row 8
	Point{0, 8}: []Point{Point{0, 7}, Point{1, 8}},
	Point{1, 8}: []Point{Point{1, 7}, Point{0, 8}, Point{2, 8}},
	Point{2, 8}: []Point{Point{2, 7}, Point{1, 8}, Point{3, 8}},
	Point{3, 8}: []Point{Point{3, 7}, Point{2, 8}, Point{4, 8}},
	Point{4, 8}: []Point{Point{4, 7}, Point{3, 8}, Point{5, 8}},
	Point{5, 8}: []Point{Point{5, 7}, Point{4, 8}, Point{4, 8}},
	Point{6, 8}: []Point{Point{6, 7}, Point{5, 8}},
}

var jumpingAdjacency = map[Point][]Point{
	// Row 0
	Point{0, 0}: []Point{Point{1, 0}, Point{0, 1}},
	Point{1, 0}: []Point{Point{0, 0}, Point{2, 0}, Point{1, 1}},
	Point{2, 0}: []Point{Point{1, 0}, Point{3, 0}, Point{2, 1}},
	Point{3, 0}: []Point{Point{2, 0}, Point{4, 0}, Point{3, 1}},
	Point{4, 0}: []Point{Point{3, 0}, Point{5, 0}, Point{4, 1}},
	Point{5, 0}: []Point{Point{4, 0}, Point{6, 0}, Point{5, 1}},
	Point{6, 0}: []Point{Point{5, 0}, Point{6, 1}},
	// Row 1
	Point{0, 1}: []Point{Point{0, 0}, Point{1, 1}, Point{0, 2}},
	Point{1, 1}: []Point{Point{1, 0}, Point{0, 1}, Point{2, 1}, Point{1, 2}},
	Point{2, 1}: []Point{Point{2, 0}, Point{1, 1}, Point{3, 1}, Point{2, 2}},
	Point{3, 1}: []Point{Point{3, 0}, Point{2, 1}, Point{4, 1}, Point{3, 2}},
	Point{4, 1}: []Point{Point{4, 0}, Point{3, 1}, Point{5, 1}, Point{4, 2}},
	Point{5, 1}: []Point{Point{5, 0}, Point{4, 1}, Point{6, 1}, Point{5, 2}},
	Point{6, 1}: []Point{Point{6, 0}, Point{5, 1}, Point{6, 2}},
	// Row 2
	Point{0, 2}: []Point{Point{0, 1}, Point{1, 2}, Point{0, 3}},
	Point{1, 2}: []Point{Point{1, 1}, Point{0, 2}, Point{2, 2}, Point{1, 6}},
	Point{2, 2}: []Point{Point{2, 1}, Point{1, 2}, Point{3, 2}, Point{2, 6}},
	Point{3, 2}: []Point{Point{3, 1}, Point{2, 2}, Point{4, 2}, Point{3, 3}},
	Point{4, 2}: []Point{Point{4, 1}, Point{3, 2}, Point{5, 2}, Point{4, 6}},
	Point{5, 2}: []Point{Point{5, 1}, Point{4, 2}, Point{6, 2}, Point{5, 6}},
	Point{6, 2}: []Point{Point{6, 1}, Point{5, 2}, Point{6, 3}},
	// Row 3
	Point{0, 3}: []Point{Point{0, 2}, Point{3, 3}, Point{0, 4}},
	Point{1, 3}: []Point{},
	Point{2, 3}: []Point{},
	Point{3, 3}: []Point{Point{3, 2}, Point{0, 3}, Point{6, 3}, Point{3, 4}},
	Point{4, 3}: []Point{},
	Point{5, 3}: []Point{},
	Point{6, 3}: []Point{Point{6, 2}, Point{3, 3}, Point{6, 4}},
	// Row 4
	Point{0, 4}: []Point{Point{0, 3}, Point{3, 4}, Point{0, 5}},
	Point{1, 4}: []Point{},
	Point{2, 4}: []Point{},
	Point{3, 4}: []Point{Point{3, 3}, Point{0, 4}, Point{6, 4}, Point{3, 5}},
	Point{4, 4}: []Point{},
	Point{5, 4}: []Point{},
	Point{6, 4}: []Point{Point{6, 3}, Point{3, 4}, Point{6, 5}},
	// Row 5
	Point{0, 5}: []Point{Point{0, 4}, Point{3, 5}, Point{0, 6}},
	Point{1, 5}: []Point{},
	Point{2, 5}: []Point{},
	Point{3, 5}: []Point{Point{3, 4}, Point{0, 5}, Point{6, 5}, Point{3, 6}},
	Point{4, 5}: []Point{},
	Point{5, 5}: []Point{},
	Point{6, 5}: []Point{Point{6, 4}, Point{3, 5}, Point{6, 6}},
	// Row 6
	Point{0, 6}: []Point{Point{0, 5}, Point{1, 6}, Point{0, 7}},
	Point{1, 6}: []Point{Point{1, 2}, Point{0, 6}, Point{2, 6}, Point{1, 7}},
	Point{2, 6}: []Point{Point{2, 2}, Point{1, 6}, Point{3, 6}, Point{2, 7}},
	Point{3, 6}: []Point{Point{3, 5}, Point{2, 6}, Point{4, 6}, Point{3, 7}},
	Point{4, 6}: []Point{Point{4, 2}, Point{3, 6}, Point{5, 6}, Point{4, 7}},
	Point{5, 6}: []Point{Point{5, 2}, Point{4, 6}, Point{6, 6}, Point{5, 7}},
	Point{6, 6}: []Point{Point{6, 5}, Point{5, 6}, Point{6, 7}},
	// Row 7
	Point{0, 7}: []Point{Point{0, 6}, Point{1, 7}, Point{0, 8}},
	Point{1, 7}: []Point{Point{1, 6}, Point{0, 7}, Point{2, 7}, Point{1, 8}},
	Point{2, 7}: []Point{Point{2, 6}, Point{1, 7}, Point{3, 7}, Point{2, 8}},
	Point{3, 7}: []Point{Point{3, 6}, Point{2, 7}, Point{4, 7}, Point{3, 8}},
	Point{4, 7}: []Point{Point{4, 6}, Point{3, 7}, Point{5, 7}, Point{4, 8}},
	Point{5, 7}: []Point{Point{5, 6}, Point{4, 7}, Point{6, 7}, Point{5, 8}},
	Point{6, 7}: []Point{Point{6, 6}, Point{5, 7}, Point{6, 8}},
	// Row 8
	Point{0, 8}: []Point{Point{0, 7}, Point{1, 8}},
	Point{1, 8}: []Point{Point{1, 7}, Point{0, 8}, Point{2, 8}},
	Point{2, 8}: []Point{Point{2, 7}, Point{1, 8}, Point{3, 8}},
	Point{3, 8}: []Point{Point{3, 7}, Point{2, 8}, Point{4, 8}},
	Point{4, 8}: []Point{Point{4, 7}, Point{3, 8}, Point{5, 8}},
	Point{5, 8}: []Point{Point{5, 7}, Point{4, 8}, Point{4, 8}},
	Point{6, 8}: []Point{Point{6, 7}, Point{5, 8}},
}

var swimmingAdjacency = map[Point][]Point{
	// Row 0
	Point{0, 0}: []Point{Point{1, 0}, Point{0, 1}},
	Point{1, 0}: []Point{Point{0, 0}, Point{2, 0}, Point{1, 1}},
	Point{2, 0}: []Point{Point{1, 0}, Point{3, 0}, Point{2, 1}},
	Point{3, 0}: []Point{Point{2, 0}, Point{4, 0}, Point{3, 1}},
	Point{4, 0}: []Point{Point{3, 0}, Point{5, 0}, Point{4, 1}},
	Point{5, 0}: []Point{Point{4, 0}, Point{6, 0}, Point{5, 1}},
	Point{6, 0}: []Point{Point{5, 0}, Point{6, 1}},
	// Row 1
	Point{0, 1}: []Point{Point{0, 0}, Point{1, 1}, Point{0, 2}},
	Point{1, 1}: []Point{Point{1, 0}, Point{0, 1}, Point{2, 1}, Point{1, 2}},
	Point{2, 1}: []Point{Point{2, 0}, Point{1, 1}, Point{3, 1}, Point{2, 2}},
	Point{3, 1}: []Point{Point{3, 0}, Point{2, 1}, Point{4, 1}, Point{3, 2}},
	Point{4, 1}: []Point{Point{4, 0}, Point{3, 1}, Point{5, 1}, Point{4, 2}},
	Point{5, 1}: []Point{Point{5, 0}, Point{4, 1}, Point{6, 1}, Point{5, 2}},
	Point{6, 1}: []Point{Point{6, 0}, Point{5, 1}, Point{6, 2}},
	// Row 2
	Point{0, 2}: []Point{Point{0, 1}, Point{1, 2}, Point{0, 3}},
	Point{1, 2}: []Point{Point{1, 1}, Point{0, 2}, Point{2, 2}, Point{1, 3}},
	Point{2, 2}: []Point{Point{2, 1}, Point{1, 2}, Point{3, 2}, Point{2, 3}},
	Point{3, 2}: []Point{Point{3, 1}, Point{2, 2}, Point{4, 2}, Point{3, 3}},
	Point{4, 2}: []Point{Point{4, 1}, Point{3, 2}, Point{5, 2}, Point{4, 3}},
	Point{5, 2}: []Point{Point{5, 1}, Point{4, 2}, Point{6, 2}, Point{5, 3}},
	Point{6, 2}: []Point{Point{6, 1}, Point{5, 2}, Point{6, 3}},
	// Row 3
	Point{0, 3}: []Point{Point{0, 2}, Point{1, 3}, Point{0, 4}},
	Point{1, 3}: []Point{Point{1, 2}, Point{0, 3}, Point{2, 3}, Point{1, 4}},
	Point{2, 3}: []Point{Point{2, 2}, Point{1, 3}, Point{3, 3}, Point{2, 4}},
	Point{3, 3}: []Point{Point{3, 2}, Point{2, 3}, Point{4, 3}, Point{3, 4}},
	Point{4, 3}: []Point{Point{4, 2}, Point{3, 3}, Point{5, 3}, Point{4, 4}},
	Point{5, 3}: []Point{Point{5, 2}, Point{4, 3}, Point{6, 3}, Point{5, 4}},
	Point{6, 3}: []Point{Point{6, 2}, Point{5, 3}, Point{6, 4}},
	// Row 4
	Point{0, 4}: []Point{Point{0, 3}, Point{1, 4}, Point{0, 5}},
	Point{1, 4}: []Point{Point{1, 3}, Point{0, 4}, Point{2, 4}, Point{1, 5}},
	Point{2, 4}: []Point{Point{2, 3}, Point{1, 4}, Point{3, 4}, Point{2, 5}},
	Point{3, 4}: []Point{Point{3, 3}, Point{2, 4}, Point{4, 4}, Point{3, 5}},
	Point{4, 4}: []Point{Point{4, 3}, Point{3, 4}, Point{5, 4}, Point{4, 5}},
	Point{5, 4}: []Point{Point{5, 3}, Point{4, 4}, Point{6, 4}, Point{5, 5}},
	Point{6, 4}: []Point{Point{6, 3}, Point{5, 4}, Point{6, 5}},
	// Row 5
	Point{0, 5}: []Point{Point{0, 4}, Point{1, 5}, Point{0, 6}},
	Point{1, 5}: []Point{Point{1, 4}, Point{0, 5}, Point{2, 5}, Point{1, 6}},
	Point{2, 5}: []Point{Point{2, 4}, Point{1, 5}, Point{3, 5}, Point{2, 6}},
	Point{3, 5}: []Point{Point{3, 4}, Point{2, 5}, Point{4, 5}, Point{3, 6}},
	Point{4, 3}: []Point{Point{4, 4}, Point{3, 5}, Point{5, 5}, Point{4, 6}},
	Point{5, 3}: []Point{Point{5, 4}, Point{4, 5}, Point{6, 5}, Point{5, 6}},
	Point{6, 3}: []Point{Point{6, 4}, Point{5, 5}, Point{6, 6}},
	// Row 6
	Point{0, 6}: []Point{Point{0, 5}, Point{1, 6}, Point{0, 7}},
	Point{1, 6}: []Point{Point{1, 5}, Point{0, 6}, Point{2, 6}, Point{1, 7}},
	Point{2, 6}: []Point{Point{2, 5}, Point{1, 6}, Point{3, 6}, Point{2, 7}},
	Point{3, 6}: []Point{Point{3, 5}, Point{2, 6}, Point{4, 6}, Point{3, 7}},
	Point{4, 6}: []Point{Point{4, 5}, Point{3, 6}, Point{5, 6}, Point{4, 7}},
	Point{5, 6}: []Point{Point{5, 5}, Point{4, 6}, Point{6, 6}, Point{5, 7}},
	Point{6, 6}: []Point{Point{6, 5}, Point{5, 6}, Point{6, 7}},
	// Row 7
	Point{0, 7}: []Point{Point{0, 6}, Point{1, 7}, Point{0, 8}},
	Point{1, 7}: []Point{Point{1, 6}, Point{0, 7}, Point{2, 7}, Point{1, 8}},
	Point{2, 7}: []Point{Point{2, 6}, Point{1, 7}, Point{3, 7}, Point{2, 8}},
	Point{3, 7}: []Point{Point{3, 6}, Point{2, 7}, Point{4, 7}, Point{3, 8}},
	Point{4, 7}: []Point{Point{4, 6}, Point{3, 7}, Point{5, 7}, Point{4, 8}},
	Point{5, 7}: []Point{Point{5, 6}, Point{4, 7}, Point{6, 7}, Point{5, 8}},
	Point{6, 7}: []Point{Point{6, 6}, Point{5, 7}, Point{6, 8}},
	// Row 8
	Point{0, 8}: []Point{Point{0, 7}, Point{1, 8}},
	Point{1, 8}: []Point{Point{1, 7}, Point{0, 8}, Point{2, 8}},
	Point{2, 8}: []Point{Point{2, 7}, Point{1, 8}, Point{3, 8}},
	Point{3, 8}: []Point{Point{3, 7}, Point{2, 8}, Point{4, 8}},
	Point{4, 8}: []Point{Point{4, 7}, Point{3, 8}, Point{5, 8}},
	Point{5, 8}: []Point{Point{5, 7}, Point{4, 8}, Point{4, 8}},
	Point{6, 8}: []Point{Point{6, 7}, Point{5, 8}},
}

func (b *Board) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("+--+--+--+--+--+--+--+\n")
	for y := 8; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			buffer.WriteString(fmt.Sprintf("|%v", b.Get(Point{x, y}).String()))
		}
		buffer.WriteString(fmt.Sprintf("|\n"))
		buffer.WriteString("+--+--+--+--+--+--+--+\n")
	}
	return buffer.String()
}

func (b *Board) Marshal() string {
	blob, _ := json.Marshal(b)
	return string(blob)
}

func Unmarshal(blob string) (*Board, error) {
	b := &Board{}
	err := json.Unmarshal([]byte(blob), b)
	return b, err
}
