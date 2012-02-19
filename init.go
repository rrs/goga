package goga

import "math/rand"

// handles problem specific initialisation takes an empty population
// and fills it with initialised *Individual
type Initialiser interface {
	Init(p Population, nGenes int)
}

// simple bool initialiser
type BoolInitialiser struct {}

func (*BoolInitialiser) Init(pop Population, nGenes int) {
	for i := 0; i < len(pop); i++ {
		// create a slice of genes and initialise them
		genes := make([]Gene, nGenes)
		for j := 0; j < nGenes; j++ {
			if rand.Intn(2) == 1 {
				genes[j] = &BoolGene{ true }
			} else {
				genes[j] = &BoolGene{ false }
			}
		}
		// construct a new *Indivdual with genes, and a fitness
		// of 0
		pop[i] = &Individual{ genes, 0 }
	}
}

func NewBoolInitialiser() Initialiser {
	return &BoolInitialiser{}
}
