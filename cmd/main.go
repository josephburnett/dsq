package main

import (
	"fmt"

	"github.com/josephburnett/dsq-golang/pkg/types"
)

func main() {
	b := types.NewBoard()
	fmt.Print(b.String())
	fmt.Printf("%v\n", b.MoveList())
}
