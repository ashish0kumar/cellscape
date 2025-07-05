package automata

import (
	"math/rand"

	"github.com/ashish0kumar/cellscape/internal/grid"
)

type ForestFire struct{}

func (ff *ForestFire) Name() string {
	return "Forest Fire"
}

func (ff *ForestFire) Description() string {
	return "Trees grow, catch fire, and burn down in cycles"
}

func (ff *ForestFire) Initialize(g *grid.Grid) {
	// Initialize with a mix of trees, fire and empty cells
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			r := rand.Float32()
			if r < 0.6 {
				g.Set(x, y, grid.Cell{Alive: true, Value: 1}) // Tree
			} else if r < 0.61 {
				g.Set(x, y, grid.Cell{Alive: true, Value: 2}) // Fire
			} else {
				g.Set(x, y, grid.Cell{Alive: false, Value: 0}) // Empty
			}
		}
	}
}

func (ff *ForestFire) Step(g *grid.Grid) *grid.Grid {
	newGrid := g.Copy()

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			cell := g.Get(x, y)

			switch cell.Value {
			case 0: // Empty -> Tree
				if rand.Float32() < 0.01 {
					newGrid.Set(x, y, grid.Cell{Alive: true, Value: 1})
				}
			case 1: // Tree
				fireNeighbors := ff.countFireNeighbors(g, x, y)
				if fireNeighbors > 0 && rand.Float32() < 0.8 {
					newGrid.Set(x, y, grid.Cell{Alive: true, Value: 2}) // Catch fire
				} else if rand.Float32() < 0.0001 {
					newGrid.Set(x, y, grid.Cell{Alive: true, Value: 2}) // Lightning
				}
			case 2: // Fire -> Empty
				newGrid.Set(x, y, grid.Cell{Alive: false, Value: 0})
			}
		}
	}

	return newGrid
}

// countFireNeighbors counts the number of fire neighbors for a cell at (x, y)
func (ff *ForestFire) countFireNeighbors(g *grid.Grid, x, y int) int {
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
