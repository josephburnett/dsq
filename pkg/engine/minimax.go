package engine

import "github.com/josephburnett/dsq-golang/pkg/types"

func BestMove(b *types.Board, s types.Side) ([2]types.Point, bool) {
	if _, m, ok := minimax(b, s, 8); ok {
		return m, true
	}
	return [2]types.Point{}, false
}

func minimax(b *types.Board, s types.Side, depth int) (int, [2]types.Point, bool) {
	if depth == 0 || b.Get(types.ADen) != types.Empty || b.Get(types.BDen) != types.Empty {
		return Fitness(b), [2]types.Point{}, false
	}
	haveMove := false
	var bestMove [2]types.Point
	var bestFitnessValue int
	for _, m := range b.MoveList() {
		if b.Get(m[0]).Side() != s {
			continue
		}
		displaced := b.Move(m)
		if s == types.A {
			min, _, _ := minimax(b, types.B, depth-1)
			// Choose the maximum of the minimums
			if min > bestFitnessValue || !haveMove {
				bestMove = m
				bestFitnessValue = min
				haveMove = true
			}
		} else {
			max, _, _ := minimax(b, types.A, depth-1)
			// Choose the minimum of the maximums
			if max < bestFitnessValue || !haveMove {
				bestMove = m
				bestFitnessValue = max
				haveMove = true
			}
		}
		b.Unmove(m, displaced)
	}
	if haveMove {
		return bestFitnessValue, bestMove, true
	}
	return 0, [2]types.Point{}, false
}
