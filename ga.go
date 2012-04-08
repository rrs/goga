package goga

// Clone: return a deep copy
// Mutate: handles gene specific mutation
// String: return a printable version of the data
type Gene interface {
	Clone() Gene
	Mutate()
	String() string
}

// handles problem specific initialisation takes an empty population
// and fills it with initialised *Individual
type Initialiser interface {
	Init(p Population, nGenes int)
}

// handles problem specific evaluation, takes a pointer to
// allow the evaluator to apply elitism
type Evaluator interface {
	Evaluate(p Population) (Population, bool)
}