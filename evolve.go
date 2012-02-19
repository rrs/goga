package goga

import "fmt"

// Function will evolve a population using the given parameters
//
// popSize: size of the population
// nGenes: the number of genes per individual
// nRuns: the maximum number of iteration before quitting
// initialiser: interface handles initialisation of the population
// evaluator: interface handles evaluation of the population
// selector: interface handles selection
// crosser: interface handles crossover
// mutater: type handles mutation

func Evolve(
	popSize, nGenes, nRuns int,
	initialiser Initialiser,
	evaluator Evaluator,
	selecter Selector,
	mater Mater,
	printPop bool) {

	// allocate population. +1 cap for elitism
	p := make(Population, popSize, popSize + 1)
	// initialise the population
	initialiser.Init(p, nGenes)
	
	// main loop
	for i := 0; i < nRuns; i++ {
		// evaluate the current population, all check for
		// termination condition, exit if condition is met
		p, exit := evaluator.Evaluate(p)
		if (exit) {
			// print the final population, if required
			if printPop {
				Print(p)
			}
			// show numer of runs to complete
			fmt.Printf("completed in %d runs", i)
			goto exit // cheeky exit :)
		}
		// select parents
		parents := selecter.Select(p)
		// crossover the parents for a new set of children
		p = mater.Mate(p, parents)
	}
	// if we are here we did not reach the termination condition
	// doh!
	evaluator.Evaluate(p)
	// print population if required
	if printPop {
		Print(p)
	}
	fmt.Printf("failed to complete after %d runs", nRuns)
	exit:
}
// prints out each individual
func Print(p Population) {
	for _,individual := range p {
		fmt.Print(individual)
	}
}
