package engine

import (
	"time"

	"github.com/josephburnett/dsq/pkg/types"
)

type Stat struct {
	Time               time.Duration
	PositionsEvaluated int
	BestOutcome        int
}

func BestMove(b *types.Board, s types.Side, depth int) ([2]types.Point, *Stat, bool) {
	start := time.Now()
	bestOutcome, bestMove, positionsEvaluated, ok := minimax(b, s, depth)
	stat := &Stat{
		Time:               time.Since(start),
		PositionsEvaluated: positionsEvaluated,
		BestOutcome:        bestOutcome,
	}
	return bestMove, stat, ok
}

func minimax(b *types.Board, s types.Side, depth int) (int, [2]types.Point, int, bool) {
	if depth == 0 || b.Get(types.ADen) != types.Empty || b.Get(types.BDen) != types.Empty {
		return Fitness(b), [2]types.Point{}, 1, false
	}
	positionsEvaluated := 0
	haveMove := false
	var bestMove [2]types.Point
	var bestFitnessValue int
	for _, m := range b.MoveList() {
		side := b.Get(m[0]).Side()
		if s != side {
			continue
		}
		displaced := b.Move(m)
		if s == types.A {
			min, _, count, _ := minimax(b, types.B, depth-1)
			positionsEvaluated += count
			// Choose the maximum of the minimums
			if min > bestFitnessValue || !haveMove {
				bestMove = m
				bestFitnessValue = min
				haveMove = true
			}
		} else {
			max, _, count, _ := minimax(b, types.A, depth-1)
			positionsEvaluated += count
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
		return bestFitnessValue, bestMove, positionsEvaluated, true
	}
	return 0, [2]types.Point{}, positionsEvaluated, false
}
