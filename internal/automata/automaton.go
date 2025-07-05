package automata

import "github.com/ashish0kumar/cellscape/internal/grid"

// Automaton defines the interface for all cellular automata
type Automaton interface {
	Step(g *grid.Grid) *grid.Grid // advances the automaton by one generation
	Initialize(g *grid.Grid)      // sets up the initial state of the grid
	Name() string                 // returns name of the automaton
	Description() string          // returns description of the automaton
}

// NewAutomaton creates a new instance of an automaton based on its name
func NewAutomaton(name string) Automaton {
	switch name {
	case "life":
		return &GameOfLife{}
	case "ant":
		return &LangtonAnt{}
	case "brain":
		return &BriansBrain{}
	case "forest":
		return &ForestFire{}
	case "wildfire":
		return &Wildfire{}
	case "belousov":
		return &BelousovZhabotinsky{}
	case "ltl":
		return &LargerThanLife{}
	case "faders":
		return &Faders{}
	default:
		return &GameOfLife{}
	}
}
