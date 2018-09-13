package types

import (
	"bytes"
	"fmt"
)

type Board [9][7]Piece
type Point [2]int

func NewBoard() *Board {
	return &Board{
		{ATiger, Empty, Empty, Empty, Empty, Empty, ALion},
		{Empty, ACat, Empty, Empty, Empty, ADog, Empty},
		{AElephant, Empty, AWolf, Empty, AHyena, Empty, AMouse},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{Empty, Empty, Empty, Empty, Empty, Empty, Empty},
		{BMouse, Empty, BHyena, Empty, BCat, Empty, BElephant},
		{Empty, BDog, Empty, Empty, Empty, BCat, Empty},
		{BLion, Empty, Empty, Empty, Empty, Empty, BTiger},
	}
}

func (b *Board) Get(p Point) Piece {
	return b[p[1]][p[0]]
}

func (b *Board) MoveList() {

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
	Point{3, 3}: []Point{Point{3, 2}, Point{0, 3}, Point{7, 3}, Point{3, 4}},
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
