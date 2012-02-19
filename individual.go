package goga

import "strconv"

// holds a slice of genes and a fitness value
// it is possible to contain any number of types of genes
type Individual struct {
	Genes []Gene
	Fitness int
}

// printable version of the individual
func (individual Individual) String() string {
	s := ""
	//print all the genes
	for _,gene := range individual.Genes {
		s += gene.String()
	}
	// followed by a space and the fitness
	s += " " + strconv.Itoa(individual.Fitness) + "\n"
	return s
}
