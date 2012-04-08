package goga

// Population has utility methods which are useful on a slice
// of type *Individual
type Population []*Individual

// removes the individual at index i and shortens the slice
func (p Population) DeleteI(i int) {
	n := len(p)
	copy(p[i:n-1], p[i+1:n])
	p = p[0:n-1]
}

// removes the individual if it exists in the Population p
func (p Population) Delete(individual  *Individual) {
	for i := range p {
		if individual  == p[i] {
			p.DeleteI(i)
			return
		}
	}
}

// returns true if individual is contained within Population p
func (p Population) Contains(individual *Individual) bool {
	for _, person := range p {
		if  person == individual {
			return true
		}
	}
	return false
}