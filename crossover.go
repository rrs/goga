package goga

import "math/rand"

// Crosser defines a cross function between two individuals 
// returns two children
type Crosser interface {
	Cross(father, mother *Individual) (*Individual, *Individual)
}

// Simple single point crossover will with probability chance randomly 
// pick a point between 1 and nGenes - 1 to cross the individuals
type SinglePointCrossover struct {
	chance float64
	nGenes int
}

func NewSinglePointCrossover(chance float64, nGenes int) Crosser {
	return &SinglePointCrossover { chance, nGenes }
}

func (spc *SinglePointCrossover) Cross(father, mother *Individual) (*Individual, *Individual) {
	kidAGenes := make([]Gene, spc.nGenes)
	kidBGenes := make([]Gene, spc.nGenes)

	// randomly pick a point along the genes with chance spc.chance
	// otherwise the crosspoint is 0 and the children are direct clones
	drawn := rand.Float64()
	crosspoint := 0
	if drawn < spc.chance {
		// add 1 and minus 2 to ensure at least one bit is crossed
		crosspoint = rand.Intn(spc.nGenes - 2) + 1
	}
	// populate the child genes with copys of the parents, swapping after
	// point crosspoint (if it was set other than 0)
	for i := 0; i < crosspoint; i++ {
		kidAGenes[i] = mother.Genes[i].Clone()
		kidBGenes[i] = father.Genes[i].Clone()
	}
	for i := crosspoint; i < spc.nGenes; i++ {
		kidAGenes[i] = father.Genes[i].Clone()
		kidBGenes[i] = mother.Genes[i].Clone()
	}
	return &Individual{kidAGenes, 0}, &Individual{kidBGenes, 0}
}

// Uniform crossover will cross each gene posistion in the two individuals
// with probability chance
type UniformCrossover struct {
	chance float64
	nGenes int
}

func NewUniformCrossover(chance float64, nGenes int) Crosser {
	return &UniformCrossover{ chance, nGenes }
}

func (uc *UniformCrossover) Cross(father, mother *Individual) (*Individual, *Individual) {
	kidAGenes := make([]Gene, uc.nGenes)
	kidBGenes := make([]Gene, uc.nGenes)
	// for all genes swap with probability uc.chance otherwise just clone
	for i := 0; i < uc.nGenes; i++ {
		drawn := rand.Float64()
		if drawn < uc.chance {
			kidAGenes[i] = mother.Genes[i].Clone()
			kidBGenes[i] = father.Genes[i].Clone()
		} else {
			kidAGenes[i] = father.Genes[i].Clone()
			kidBGenes[i] = mother.Genes[i].Clone()
		}
	}
	return &Individual{kidAGenes, 0}, &Individual{kidBGenes, 0}
}
