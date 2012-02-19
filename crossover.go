package goga

import "math/rand"

type Crosser interface {
	Cross(father, mother *Individual) (*Individual, *Individual)
}

type SinglePointCrossover struct {
	chance float64	
	nGenes int
}

func NewSinglePointCrossover(chance float64, nGenes int) Crosser {
	return &SinglePointCrossover {chance, nGenes}
}

func (spc *SinglePointCrossover) Cross(father, mother *Individual) (*Individual,*Individual) {	
	kidAGenes := make([]Gene, spc.nGenes)
	kidBGenes := make([]Gene, spc.nGenes)

	drawn := rand.Float64()
	crosspoint := 0
	if drawn < spc.chance {
		crosspoint = rand.Intn(spc.nGenes) 
	}

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