package engine

import (
	"github.com/josephburnett/dsq-golang/pkg/types"
)

func Fitness(b *types.Board) int {
	fitness := 0
	for x := 0; x < 7; x++ {
		for y := 0; y < 9; y++ {
			pt := types.Point{x, y}
			piece := b.Get(pt)
			fitness += materialValue(piece)
			fitness += developmentValue(piece, pt)
		}
	}
	return fitness
}

func materialValue(p types.Piece) int {
	switch p {
	case types.AMouse:
		return 500
	case types.ACat:
		return 200
	case types.AWolf:
		return 300
	case types.ADog:
		return 400
	case types.AHyena:
		return 500
	case types.ATiger:
		return 800
	case types.ALion:
		return 900
	case types.AElephant:
		return 1000
	case types.BMouse:
		return -500
	case types.BCat:
		return -200
	case types.BWolf:
		return -300
	case types.BDog:
		return -400
	case types.BHyena:
		return -500
	case types.BTiger:
		return -800
	case types.BLion:
		return -900
	case types.BElephant:
		return -1000
	default:
		return 0
	}
}

func developmentValue(p types.Piece, pt types.Point) int {
	switch p {
	case types.AMouse:
		return mouseDevelopment.Get(pt)
	case types.ACat:
		return catDevelopment.Get(pt)
	case types.AWolf:
		return wolfDevelopment.Get(pt)
	case types.ADog:
		return dogDevelopment.Get(pt)
	case types.AHyena:
		return hyenaDevelopment.Get(pt)
	case types.ATiger:
		return tigerDevelopment.Get(pt)
	case types.ALion:
		return lionDevelopment.Get(pt)
	case types.AElephant:
		return elephantDevelopment.Get(pt)
	case types.BMouse:
		return -mouseDevelopment.Get(pt.Rotate())
	case types.BCat:
		return -catDevelopment.Get(pt.Rotate())
	case types.BWolf:
		return -wolfDevelopment.Get(pt.Rotate())
	case types.BDog:
		return -dogDevelopment.Get(pt.Rotate())
	case types.BHyena:
		return -hyenaDevelopment.Get(pt.Rotate())
	case types.BTiger:
		return -tigerDevelopment.Get(pt.Rotate())
	case types.BLion:
		return -lionDevelopment.Get(pt.Rotate())
	case types.BElephant:
		return -elephantDevelopment.Get(pt.Rotate())
	default:
		return 0
	}
}

const win = 9999

type development [9][7]int

func (d development) Get(p types.Point) int {
	return d[p[1]][p[0]]
}

var mouseDevelopment = development{
	{8, 8, 8, 0, 8, 8, 8},
	{8, 8, 8, 9, 9, 9, 9},
	{8, 8, 8, 9, 10, 10, 10},
	{8, 9, 9, 10, 12, 12, 11},
	{8, 9, 9, 11, 12, 12, 12},
	{8, 9, 9, 11, 12, 12, 13},
	{10, 11, 11, 13, 13, 13, 13},
	{11, 12, 13, 50, 13, 13, 13},
	{11, 13, 50, win, 50, 13, 13},
}

var catDevelopment = development{
	{8, 8, 8, 0, 8, 8, 8},
	{13, 10, 8, 8, 8, 8, 8},
	{10, 10, 10, 8, 8, 8, 8},
	{10, 0, 0, 8, 0, 0, 8},
	{10, 0, 0, 8, 0, 0, 8},
	{10, 0, 0, 10, 0, 0, 8},
	{10, 11, 11, 15, 11, 11, 10},
	{11, 11, 15, 50, 15, 11, 11},
	{11, 15, 50, win, 50, 15, 11},
}

var wolfDevelopment = development{
	{8, 12, 12, 0, 8, 8, 8},
	{8, 12, 13, 8, 8, 8, 8},
	{8, 8, 10, 8, 8, 8, 8},
	{8, 0, 0, 8, 0, 0, 8},
	{8, 0, 0, 8, 0, 0, 8},
	{9, 0, 0, 10, 0, 0, 9},
	{9, 10, 11, 15, 11, 10, 9},
	{10, 11, 15, 50, 15, 11, 10},
	{11, 15, 50, win, 50, 15, 11},
}

var dogDevelopment = development{
	{8, 8, 8, 0, 12, 12, 8},
	{8, 8, 8, 8, 13, 10, 8},
	{8, 8, 8, 8, 8, 8, 8},
	{8, 0, 0, 8, 0, 0, 8},
	{8, 0, 0, 8, 0, 0, 8},
	{9, 0, 0, 10, 0, 0, 9},
	{10, 11, 15, 50, 15, 11, 10},
	{11, 15, 50, win, 50, 15, 11},
}

var hyenaDevelopment = development{
	{9, 9, 9, 0, 9, 9, 9},
	{9, 9, 9, 9, 9, 9, 9},
	{9, 9, 9, 10, 10, 9, 9},
	{10, 0, 0, 13, 0, 0, 10},
	{11, 0, 0, 14, 0, 0, 11},
	{12, 0, 0, 15, 0, 0, 12},
	{13, 13, 14, 15, 14, 13, 13},
	{13, 14, 15, 50, 15, 14, 13},
	{14, 15, 50, win, 50, 15, 14},
}

var tigerDevelopment = development{
	{10, 12, 12, 0, 12, 12, 10},
	{12, 14, 12, 12, 12, 12, 12},
	{14, 16, 16, 14, 16, 16, 14},
	{15, 0, 0, 15, 0, 0, 15},
	{15, 0, 0, 15, 0, 0, 15},
	{15, 0, 0, 15, 0, 0, 15},
	{18, 20, 20, 30, 20, 20, 18},
	{25, 25, 30, 50, 30, 25, 25},
	{25, 30, 50, win, 50, 30, 25},
}

var lionDevelopment = development{
	{10, 12, 12, 0, 12, 12, 10},
	{12, 14, 12, 12, 12, 12, 12},
	{14, 16, 16, 14, 16, 16, 14},
	{15, 0, 0, 15, 0, 0, 15},
	{15, 0, 0, 15, 0, 0, 15},
	{15, 0, 0, 15, 0, 0, 15},
	{18, 20, 20, 30, 20, 20, 18},
	{25, 25, 30, 50, 30, 25, 25},
	{25, 30, 50, win, 50, 30, 25},
}

var elephantDevelopment = development{
	{11, 11, 11, 0, 11, 11, 11},
	{11, 11, 11, 11, 11, 11, 11},
	{10, 15, 14, 14, 14, 14, 12},
	{12, 0, 0, 12, 0, 0, 12},
	{14, 0, 0, 14, 0, 0, 14},
	{16, 0, 0, 16, 0, 0, 16},
	{18, 20, 20, 30, 20, 20, 18},
	{25, 25, 30, 50, 30, 25, 25},
	{25, 30, 50, win, 50, 30, 25},
}
