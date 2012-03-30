package main

import (
	//"fmt"
	"runtime"
	"math/rand"
	"time"
	"goga"
)

const (
	POP_SIZE = 128
	N_PARENTS = POP_SIZE / 2
	N_GENES = 512
	TOURNAMENT_SIZE = 3
	Cp = 0.3
	Mp = 0.003
	N_RUNS = 1000
)


func main() {
	runtime.GOMAXPROCS(8)
	t := time.Now()
	rand.Seed(t.Unix())

	initialiser := goga.NewBoolInitialiser()
	evaluator := goga.NewBoolEvaluator(N_GENES)
	//selector := goga.NewRouletteSelection(N_PARENTS)
	selector := goga.NewTournamentSelection(N_PARENTS, TOURNAMENT_SIZE)
	crosser := goga.NewUniformCrossover(Cp, N_GENES)
	mutator := goga.NewSimpleMutator(Mp)
	mater := goga.NewProceduralMater(crosser, mutator)
	//mater := goga.NewParallelMater(crosser, mutator)
	goga.Evolve(
		POP_SIZE, N_GENES, N_RUNS,
		initialiser,
		evaluator,
		selector,
		mater,
		false,
	)
}
