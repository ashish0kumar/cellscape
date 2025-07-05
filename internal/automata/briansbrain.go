package automata

import (
	"math/rand"

	"github.com/ashish0kumar/cellscape/internal/grid"
)

type BriansBrain struct{}

func (bb *BriansBrain) Name() string {
	return "Brian's Brain"
}

func (bb *BriansBrain) Description() string {
	return "A 3-state automaton with live, dying, and dead cells"
}

func (bb *BriansBrain) Initialize(g *grid.Grid) {
	// Initialize with 25% live cells
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if rand.Float32() < 0.25 {
				g.Set(x, y, grid.Cell{Alive: true, Value: 1}) // Live = 1
			} else {
				g.Set(x, y, grid.Cell{Alive: false, Value: 0}) // Dead = 0
			}
		}
	}
}

func (bb *BriansBrain) Step(g *grid.Grid) *grid.Grid {
	newGrid := g.Copy()

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			cell := g.Get(x, y)
			liveNeighbors := bb.countLiveNeighbors(g, x, y)

			switch cell.Value {
			case 1: // Live -> Dying
				newGrid.Set(x, y, grid.Cell{Alive: true, Value: 2})
			case 2: // Dying -> Dead
				newGrid.Set(x, y, grid.Cell{Alive: false, Value: 0})
			case 0: // Dead -> Live if exactly 2 live neighbors
				if liveNeighbors == 2 {
					newGrid.Set(x, y, grid.Cell{Alive: true, Value: 1})
				}
			}
		}
	}

	return newGrid
}

// countLiveNeighbors counts the number of live neighbors for a cell at (x, y)
func (bb *BriansBrain) countLiveNeighbors(g *grid.Grid, x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height {
				if g.Get(nx, ny).Value == 1 { // Only count live cells
					count++
				}
			}
		}
	}
	return count
}
