package main

import (
	"fmt"

	"github.com/josephburnett/dsq-golang/pkg/engine"
	"github.com/josephburnett/dsq-golang/pkg/types"
)

func main() {
	b := types.NewBoard()
	fmt.Print(b.String())
	fmt.Printf("Moves: %v\n", b.MoveList())
	fmt.Printf("Fitness: %v\n", engine.Fitness(b))
	if m, _, ok := engine.BestMove(b, types.A, 4); ok {
		fmt.Printf("BestMove: %v\n", m)
	}
}
