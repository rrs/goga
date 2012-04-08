package main

import(
	"goga"
	"math/rand"
)

// simple bool initialiser
type BoolInitialiser struct {}

func (*BoolInitialiser) Init(pop goga.Population, nGenes int) {
	for i := 0; i < len(pop); i++ {
		// create a slice of genes and initialise them
		genes := make([]goga.Gene, nGenes)
		for j := 0; j < nGenes; j++ {
			if rand.Intn(2) == 1 {
				genes[j] = &BoolGene{ true }
			} else {
				genes[j] = &BoolGene{ false }
			}
		}
		// construct a new *Indivdual with genes, and a fitness
		// of 0
		pop[i] = &goga.Individual{ genes, 0 }
	}
}

func NewBoolInitialiser() goga.Initialiser {
	return &BoolInitialiser{}
}
