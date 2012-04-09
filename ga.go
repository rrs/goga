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

// handles problem specific evaluation returns a bool value to alert
// the caller to succesfully achieving the termination condition and
// retuns Population to allow the use of elitism (see allones example)
type Evaluator interface {
	Evaluate(p Population) (Population, bool)
}
