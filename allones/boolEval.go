package main

import  "goga"

//simple bool evaluator, applies elitism
type BoolEvaluator struct {
	best goga.Individual
	maxValue int
}

func (e *BoolEvaluator) Evaluate(p goga.Population) (goga.Population, bool) {
	// start by setting the best individual we have found to
	// the first individual
	best := p[0]
	popSize := len(p)
	if popSize % 2 > 0 {
		p = p[:len(p) - 1]
		popSize -= 1
	}
	for _, individual := range p {
		fitness := 0
		// total up the value of all the genes
		for _, gene := range individual.Genes {
			if gene.(*BoolGene).Data {
				fitness++
			}
		}
		// set the fitness on the current individual
		individual.Fitness = fitness
		// check if it has beaten our best in this
		// population so far as set it accordingly
		if individual.Fitness > best.Fitness {
			best = individual
		}
	}
	// once finished check if the best of this population
	// has met the termination criteria, in this case its
	// easy, let he caller know if we have 
	if (best.Fitness == e.maxValue) {
		return p, true
	}
	// check if the best fitness from this population
	// is worse than our overall best so far and re-insert
	// the overall best, otherwise check if its better and
	// set it to be the new best
	if best.Fitness < e.best.Fitness {
		p = p[0:popSize + 1]
		p[popSize] = &e.best
	} else if best.Fitness > e.best.Fitness {
		e.best = *best
	}

	return p, false
}

func NewBoolEvaluator(maxValue int) goga.Evaluator {
	return &BoolEvaluator{  maxValue:maxValue }
}
