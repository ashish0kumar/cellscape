package automata

import (
	"math/rand"

	"github.com/ashish0kumar/cellscape/internal/grid"
)

type Faders struct{}

func (f *Faders) Name() string {
	return "Faders"
}

func (f *Faders) Description() string {
	return "Cells fade through multiple states before dying"
}

func (f *Faders) Initialize(g *grid.Grid) {
	// Initialize with some random cells
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if rand.Float32() < 0.001 {
				g.Set(x, y, grid.Cell{Alive: true, Value: 1})
			}
		}
	}
}

func (f *Faders) Step(g *grid.Grid) *grid.Grid {
	newGrid := g.Copy()

	N := 127     // Max fade level
	L, U := 2, 2 // Birth range
	K, Y := 2, 2 // Survival range

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			oldState := g.Get(x, y).Value
			sum8 := f.sumNeighbors(g, x, y)

			var newState int

			if oldState == 0 && sum8 >= L && sum8 <= U {
				newState = 1
			} else if oldState == 1 {
				if sum8 >= K && sum8 <= Y {
					newState = 1
				} else {
					newState = 2
				}
			} else if (oldState&1) == 0 && oldState > 0 && oldState < (2*N) {
				newState = oldState + 2
			}

			newGrid.Set(x, y, grid.Cell{
				Alive: newState > 0,
				Value: newState,
			})
		}
	}

	return newGrid
}

// sumNeighbors calculates the sum of the values of the 8 neighbors
func (f *Faders) sumNeighbors(g *grid.Grid, x, y int) int {
	sum := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height {
				sum += g.Get(nx, ny).Value
			}
		}
	}
	return sum
}
