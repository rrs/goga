package goga

import "math/rand"

// A Selecter is used for the selection process, select returns a group
// of individuals selected by some probabilistic method giving a higher
// chance to individuals with a higher fitness
type Selector interface {
	Select(population Population) Population
}

// Tounament selection selects a group of individuals of size tournamentSize
// completely at random from the population, then the fittest individual is
// selected from that group, the lower the group size the higher the selection
// pressure on weak individuals, this is often desirable.
type TournamentSelection struct {
	nParents int
	tounamentSize int
}

func NewTournamentSelection(nParents, tounamentSize int) Selector {
	return &TournamentSelection{ nParents, tounamentSize }
}

func (ts* TournamentSelection) Select(pop Population) Population {
	// create a slice for the parents, we will just point at the
	// selected individuals
	selected := make(Population, ts.nParents)
	// keep looping until the selected slice is full
	for i := 0; i < ts.nParents; {
		// do one round of the tournament
		winner := ts.tournament(pop)
		// check if the winner is already in the selection pool
		if selected.Contains(winner) {
			continue
		}
		// if not add the winner
		selected[i] = winner
		i++
	}
	return selected
}

// Does a round in the tounament and returns the winner
func (ts* TournamentSelection) tournament(pop Population) *Individual{

	competitors := make(Population, ts.tounamentSize)
	// keep looping till the slice is full
	for i := 0; i < ts.tounamentSize; {
		// pick one at random
		index := rand.Intn(len(pop))
		// check if we already have it in the tournament
		if competitors.Contains(pop[index]) {
			continue
		}
		// if not add it
		competitors[i] = pop[index]
		i++
	}
	// loop over the competitors and find the best
	best := competitors[0]
	for i := 1; i < ts.tounamentSize; i++ {
		if competitors[i].Fitness > best.Fitness {
			best = competitors[i]
		}
	}
	return best
}

type RouletteSelection struct {
	nParents int
}

func NewRouletteSelection(nParents int) Selector {
	return &RouletteSelection { nParents }
}

var fitnessSum int

func (rs *RouletteSelection) Select(pop Population) Population {
	selected := make(Population, rs.nParents)
	fitnessSum = 0
	for _, individual := range pop {
		fitnessSum += individual.Fitness
	}

	for i := 0; i < len(selected); {
		winner := rs.round(pop)
		if !selected.Contains(winner) {
			selected[i] = winner
			i++
		}
	}
	return selected
}
func (rs *RouletteSelection) round(pop Population) *Individual {
	sum := 0
	i := 0
	drawn := rand.Intn(fitnessSum)
	for sum <= drawn {
		sum += pop[i].Fitness
		i++
	}
	return pop[i-1]
}