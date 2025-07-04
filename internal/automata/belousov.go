package automata

import (
	"cellscape/internal/grid"
	"math/rand"
)

type BelousovZhabotinsky struct{}

func (bz *BelousovZhabotinsky) Name() string {
	return "Belousov-Zhabotinsky"
}

func (bz *BelousovZhabotinsky) Description() string {
	return "Chemical reaction creating spiral waves and patterns"
}

func (bz *BelousovZhabotinsky) Initialize(g *grid.Grid) {
	// Initialize with some infected cells
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if rand.Float32() < 0.001 {
				g.Set(x, y, grid.Cell{Alive: true, Value: 1})
			} else {
				g.Set(x, y, grid.Cell{Alive: false, Value: 0})
			}
		}
	}
}

func (bz *BelousovZhabotinsky) Step(g *grid.Grid) *grid.Grid {
	newGrid := g.Copy()

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			newValue := bz.calculateIllScore(g, x, y)

			if newValue > 255 {
				newValue = 255
			}

			newGrid.Set(x, y, grid.Cell{
				Alive: newValue > 0,
				Value: newValue,
			})
		}
	}

	return newGrid
}

// calculateIllScore computes the new state based on the B-Z rules
func (bz *BelousovZhabotinsky) calculateIllScore(g *grid.Grid, x, y int) int {
	state := g.Get(x, y).Value
	if state >= 255 {
		return 0
	}

	a, b, s := 0.0, 0.0, 0.0

	// Count neighbors
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height {
				lvl := g.Get(nx, ny).Value
				s += float64(lvl)
				if lvl > 0 && lvl < 255 {
					a += 1.0
				} else if lvl == 255 {
					b += 1.0
				}
			}
		}
	}

	var newState int
	if state <= 0 {
		newState = int(a/1) + int(b/1)
	} else {
		newState = int(s/(a+b+1)) + 30
	}

	return newState
}
