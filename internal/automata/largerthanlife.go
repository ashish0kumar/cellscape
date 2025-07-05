package automata

import (
	"math/rand"

	"github.com/ashish0kumar/cellscape/internal/grid"
)

type LargerThanLife struct{}

func (ltl *LargerThanLife) Name() string {
	return "Larger than Life"
}

func (ltl *LargerThanLife) Description() string {
	return "Generalized Game of Life with larger neighborhoods"
}

func (ltl *LargerThanLife) Initialize(g *grid.Grid) {
	// Initialize with 25% live cells
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if rand.Float32() < 0.25 {
				g.Set(x, y, grid.Cell{Alive: true, Value: 1})
			}
		}
	}
}

func (ltl *LargerThanLife) Step(g *grid.Grid) *grid.Grid {
	newGrid := g.Copy()

	R := 5           // Neighborhood radius
	B1, B2 := 34, 45 // Birth range
	S1, S2 := 34, 58 // Survival range

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			neighbors := ltl.countNeighbors(g, x, y, R)
			isAlive := g.Get(x, y).Alive

			if isAlive && neighbors >= S1 && neighbors <= S2 {
				newGrid.Set(x, y, grid.Cell{Alive: true, Value: 1})
			} else if neighbors >= B1 && neighbors <= B2 {
				newGrid.Set(x, y, grid.Cell{Alive: true, Value: 1})
			} else {
				newGrid.Set(x, y, grid.Cell{Alive: false, Value: 0})
			}
		}
	}

	return newGrid
}

// countNeighbors counts the number of alive neighbors within a given radius
func (ltl *LargerThanLife) countNeighbors(g *grid.Grid, x, y, radius int) int {
	count := 0
	for dy := -radius; dy <= radius; dy++ {
		for dx := -radius; dx <= radius; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height {
				if g.Get(nx, ny).Alive {
					count++
				}
			}
		}
	}
	return count
}
