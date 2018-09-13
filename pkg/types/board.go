package types

import (
	"bytes"
	"fmt"
)

type Board [9][7]Piece

func NewBoard() *Board {
	return &Board{
		{
			ATiger,
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
			ALion,
		},
		{
			Empty,
			ACat,
			Empty,
			Empty,
			Empty,
			ADog,
			Empty,
		},
		{
			AElephant,
			Empty,
			AWolf,
			Empty,
			AHyena,
			Empty,
			AMouse,
		},
		{
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
		},
		{
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
		},
		{
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
		},
		{
			BMouse,
			Empty,
			BHyena,
			Empty,
			BCat,
			Empty,
			BElephant,
		},
		{
			Empty,
			BDog,
			Empty,
			Empty,
			Empty,
			BCat,
			Empty,
		},
		{
			BLion,
			Empty,
			Empty,
			Empty,
			Empty,
			Empty,
			BTiger,
		},
	}
}

func (b *Board) ADen() Piece {
	return b[3][6]
}

func (b *Board) BDen() Piece {
	return b[3][0]
}

func (b *Board) Get(x, y int) Piece {
	return b[x][y]
}

func (b *Board) CanMove(fromX, fromY, toX, toY int) bool {
	p := b.Get(fromX, fromY)
	if p == Empty {
		return false
	}
	return false
}

func (b *Board) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("+--+--+--+--+--+--+--+\n")
	for y := 8; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			buffer.WriteString(fmt.Sprintf("|%v", b[y][x].String()))
		}
		buffer.WriteString(fmt.Sprintf("|\n"))
		buffer.WriteString("+--+--+--+--+--+--+--+\n")
	}
	return buffer.String()
}
