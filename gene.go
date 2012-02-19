package goga

// defines what is required of a gene
// Clone: must return a deep copy
// Mutate: handles gene specific mutation
// String: return a printable version of the data
type Gene interface {
	Clone() Gene
	Mutate()
	String() string
}

// simple bool type gene
type BoolGene struct {
	Data bool;
}

func (gene *BoolGene) Clone() Gene {
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
