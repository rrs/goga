package goga

// handles problem specific evaluation, takes a pointer to
// allow the evaluator to apply elitism
type Evaluator interface {
	Evaluate(p Population) (Population, bool)
}

//simple bool evaluator, applies elitism
type BoolEvaluator struct {
	best Individual
	maxValue int
}

func (evaluator *BoolEvaluator) Evaluate(p Population) (Population, bool) {
	// start by setting the best individual we have found to
	// the first individual
	best := p[0]
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
	if (best.Fitness == evaluator.maxValue) {
		return p, true
	}
	// check if the best fitness from this population
	// is worse than our overall best so far and re-insert
	// the overall best, otherwise check if its better and
	// set it to be the new best
	popSize := len(p)
	if best.Fitness < evaluator.best.Fitness {
		p = p[0:popSize + 1]
		p[popSize] = &evaluator.best
	} else if best.Fitness > evaluator.best.Fitness {
		evaluator.best = *best
	}

	return p, false
}

func NewBoolEvaluator(maxValue int) Evaluator {
	return &BoolEvaluator{  maxValue:maxValue }
}
