package goga

// Combines the  crossover and mutation steps
type Mater interface {
	Mate(Population, Population) Population
}

// uses crossover and mutation in a procedural manner
type ProceduralMater struct {
	crosser Crosser
	mutator Mutator
}

func NewProceduralMater(crosser Crosser, mutator Mutator) Mater {
	return &ProceduralMater{ crosser, mutator }
}

func (mp *ProceduralMater) Mate(pop, parents Population) Population {
	// copy the parents to the end half of the slice
	copy(pop[len(parents):], parents)
	// crossover parents in pairs and mutate the offspring placing them at
	// the first half of the population
	for i := 0; i < len(parents); i+=2 {
		pop[i], pop[i+1] = mp.crosser.Cross(parents[i], parents[i+1])
		mp.mutator.Mutate(pop[i])
		mp.mutator.Mutate(pop[i+1])
	}
	return pop
}

// uses crossover and mutation in a parallel manner
type ParallelMater struct {
	crosser Crosser
	mutator Mutator
}

func NewParallelMater(crosser Crosser, mutator Mutator) Mater {
	return &ParallelMater{ crosser, mutator }
}

func (mp *ParallelMater) Mate(pop, parents Population) Population {
	// prep a channel for returning offspring
	ch := make(chan *Individual)
	// for each crossover and mutation pair create a new goroutine
	for i := 0; i < len(parents); i+=2 {
		go func(n int) {
			kidA, kidB := mp.crosser.Cross(parents[n], parents[n+1])
			mp.mutator.Mutate(kidA)
			mp.mutator.Mutate(kidB)
			ch<- kidA
			ch<- kidB
		}(i)
	}
	// while the routines finish copy the parents to the last half of the population
	copy(pop[len(parents):], parents)
	// reign in the results from the routines, placing them at the first half of population
	for i := 0; i < len(parents); i++ {
		pop[i] = <-ch
	}

	return pop
}
