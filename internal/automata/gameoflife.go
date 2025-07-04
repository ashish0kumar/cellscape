package automata

import (
	"cellscape/internal/grid"
	"math/rand"
)

type GameOfLife struct{}

func (gol *GameOfLife) Name() string {
	return "Conway's Game of Life"
}

func (gol *GameOfLife) Description() string {
	return "The classic cellular automaton with simple rules that create complex patterns"
}

func (gol *GameOfLife) Initialize(g *grid.Grid) {
	// Create a random initial state with 30% live cells
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if rand.Float32() < 0.3 {
				g.SetAlive(x, y, true)
			}
		}
	}
}

func (gol *GameOfLife) Step(g *grid.Grid) *grid.Grid {
	newGrid := g.Copy()

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			neighbors := g.CountLiveNeighbors(x, y)
			isAlive := g.IsAlive(x, y)

			// Apply rules
			if isAlive {
				// Live cell with 2 or 3 neighbors survives
				newGrid.SetAlive(x, y, neighbors == 2 || neighbors == 3)
			} else {
				// Dead cell with exactly 3 neighbors becomes alive
				newGrid.SetAlive(x, y, neighbors == 3)
			}
		}
	}

	return newGrid
}
