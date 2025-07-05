package automata

import (
	"github.com/ashish0kumar/cellscape/internal/grid"
)

// Direction represents the ant's facing direction
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type LangtonAnt struct {
	X, Y      int
	Direction Direction
}

func (la *LangtonAnt) Name() string {
	return "Langton's Ant"
}

func (la *LangtonAnt) Description() string {
	return "A simple ant that follows two rules: turn left on white, turn right on black"
}

func (la *LangtonAnt) Initialize(g *grid.Grid) {
	// Clear the grid and place the ant in the center
	g.Clear()
	la.X = g.Width / 2
	la.Y = g.Height / 2
	la.Direction = North
}

func (la *LangtonAnt) Step(g *grid.Grid) *grid.Grid {
	newGrid := g.Copy()

	// Check current cell
	currentCell := g.IsAlive(la.X, la.Y)

	if currentCell {
		// black cell: turn right, flip cell to white, move forward
		la.turnRight()
		newGrid.SetAlive(la.X, la.Y, false)
	} else {
		// white cell: turn left, flip cell to black, move forward
		la.turnLeft()
		newGrid.SetAlive(la.X, la.Y, true)
	}

	// Move forward
	la.moveForward(g.Width, g.Height)

	return newGrid
}

func (la *LangtonAnt) turnLeft() {
	la.Direction = Direction((int(la.Direction) + 3) % 4)
}

func (la *LangtonAnt) turnRight() {
	la.Direction = Direction((int(la.Direction) + 1) % 4)
}

func (la *LangtonAnt) moveForward(width, height int) {
	switch la.Direction {
	case North:
		la.Y = (la.Y - 1 + height) % height
	case East:
		la.X = (la.X + 1) % width
	case South:
		la.Y = (la.Y + 1) % height
	case West:
		la.X = (la.X - 1 + width) % width
	}
}
