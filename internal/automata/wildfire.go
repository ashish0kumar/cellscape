package automata

import (
	"cellscape/internal/grid"
	"math/rand"
)

type Wildfire struct{}

func (wf *Wildfire) Name() string {
	return "Wildfire"
}

func (wf *Wildfire) Description() string {
	return "Stochastic wildfire spread through vegetation"
}

func (wf *Wildfire) Initialize(g *grid.Grid) {
	// Initialize with a mix of vegetation, fire, and empty cells
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			r := rand.Float32()
			if r < 0.95 {
				g.Set(x, y, grid.Cell{Alive: true, Value: 1}) // Vegetation
			} else if r < 0.955 {
				g.Set(x, y, grid.Cell{Alive: true, Value: 2}) // Fire
			} else {
				g.Set(x, y, grid.Cell{Alive: false, Value: 0}) // Empty
			}
		}
	}
}

func (wf *Wildfire) Step(g *grid.Grid) *grid.Grid {
	newGrid := g.Copy()

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			cell := g.Get(x, y)

			switch cell.Value {
			case 1: // Vegetation
				fireNeighbors := wf.countFireNeighbors(g, x, y)
				igniteChance := float32(fireNeighbors) * 0.15
				if rand.Float32() < igniteChance {
					newGrid.Set(x, y, grid.Cell{Alive: true, Value: 2})
				}
			case 2: // Fire -> Burnt
				if rand.Float32() < 0.1 {
					newGrid.Set(x, y, grid.Cell{Alive: false, Value: 3})
				}
			}
		}
	}

	return newGrid
}

// countFireNeighbors counts the number of fire neighbors for a cell at (x, y)
func (wf *Wildfire) countFireNeighbors(g *grid.Grid, x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height {
				if g.Get(nx, ny).Value == 2 {
					count++
				}
			}
		}
	}
	return count
}
