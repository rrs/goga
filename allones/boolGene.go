package main

import "goga"

// simple bool type gene
type BoolGene struct {
	Data bool;
}

func (gene *BoolGene) Clone() goga.Gene {
	return &BoolGene { gene.Data }
}

func (gene *BoolGene) String() string {
	if gene.Data {
		return "1"
	}
	return "0"
}

// simply flip the bool
func (gene *BoolGene) Mutate() {
	gene.Data = !gene.Data
}