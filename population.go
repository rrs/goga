package goga

// Population has methods which are useful for use on a slice
// of *Individual
type Population []*Individual

// removes the individual at index i and shortens the slice
func (p Population) delete(i int) {
	n := len(p)
	copy(p[i:n-1], p[i+1:n])
	p = p[0:n-1]
}
// removes the individual indi if it exists from the Population p
func (p Population) Delete(individual  *Individual) {
	for i := range p {
		if individual  == p[i] {
			p.delete(i)
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
