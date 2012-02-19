package goga

type Mater interface {
	Mate(Population, Population) Population
}

type ProceduralMater struct {
	crosser Crosser
	mutator Mutator
}

func NewProceduralMater(crosser Crosser, mutator Mutator) Mater {
	return &ProceduralMater{ crosser, mutator }
}

func (mp *ProceduralMater) Mate(pop, parents Population) Population {
	copy(pop[len(parents):len(pop)], parents)
	for i := 0; i < len(parents); i+=2 {
		pop[i], pop[i+1] = mp.crosser.Cross(parents[i], parents[i+1])
		mp.mutator.Mutate(pop[i])
		mp.mutator.Mutate(pop[i+1])
	}
	return pop
}

type ParallelMater struct {
	crosser Crosser
	mutator Mutator
}

func NewParallelMater(crosser Crosser, mutator Mutator) Mater {
	return &ParallelMater{ crosser, mutator }
}

func (mp *ParallelMater) Mate(pop, parents Population) Population {
	ch := make(chan *Individual)
	for i := 0; i < len(parents); i+=2 {
		go func(n int) {
			kidA, kidB := mp.crosser.Cross(parents[n], parents[n+1])
			mp.mutator.Mutate(kidA)
			mp.mutator.Mutate(kidB)
			ch<- kidA
			ch<- kidB
		}(i)
	}
	pop = pop[0:len(parents)*2]
	copy(pop[len(parents):len(pop)], parents)
	for i := 0; i < len(parents); i++ {
		pop[i] = <-ch
	}

	return pop
}