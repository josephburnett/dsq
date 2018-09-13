package main

import (
	"fmt"

	"github.com/josephburnett/dsq-golang/pkg/types"
)

func main() {
	fmt.Print(types.NewBoard().String())
}
