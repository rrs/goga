package goga

import "math/rand"

type Mutator interface {
	Mutate(*Individual)
}

type SimpleMutator struct {
	chance float64
}

func NewSimpleMutator(chance float64) Mutator {
	return &SimpleMutator{ chance }
}

func (sm *SimpleMutator) Mutate(individual *Individual) {	
	genes := individual.Genes // []Gene
	for i := 0; i < len(genes); i++ {
		drawn := rand.Float64()
		if (drawn < sm.chance) {
			genes[i].Mutate()
		}
	}
}