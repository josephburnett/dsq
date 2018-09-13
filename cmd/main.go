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
}
