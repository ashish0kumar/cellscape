package grid

// Cell represents a single cell in the grid
type Cell struct {
	Alive bool
	Value int
}

// Grid represents the cellular automaton grid
type Grid struct {
	Width  int
	Height int
	Cells  [][]Cell
}

// NewGrid creates a new grid with the specified dimensions
func NewGrid(width, height int) *Grid {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}

	return &Grid{
		Width:  width,
		Height: height,
		Cells:  cells,
	}
}

// Get returns the cell at the specified coords
func (g *Grid) Get(x, y int) Cell {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return Cell{Alive: false, Value: 0}
	}
	return g.Cells[y][x]
}

// Set sets the cell at the specified coords
func (g *Grid) Set(x, y int, cell Cell) {
	if x >= 0 && x < g.Width && y >= 0 && y < g.Height {
		g.Cells[y][x] = cell
	}
}

// IsAlive checks if the cell at the specified coords is alive
func (g *Grid) IsAlive(x, y int) bool {
	return g.Get(x, y).Alive
}

// SetAlive sets the alive status of the cell at the specified coords
func (g *Grid) SetAlive(x, y int, alive bool) {
	if x >= 0 && x < g.Width && y >= 0 && y < g.Height {
		g.Cells[y][x].Alive = alive
	}
}

// CountLiveNeighbors counts the number of live neighbors for a cell
func (g *Grid) CountLiveNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if g.IsAlive(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

// Clear resets all cells to empty
func (g *Grid) Clear() {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.Cells[y][x] = Cell{Alive: false, Value: 0}
		}
	}
}

// Copy creates a deep copy of the grid
func (g *Grid) Copy() *Grid {
	newGrid := NewGrid(g.Width, g.Height)
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			newGrid.Cells[y][x] = g.Cells[y][x]
		}
	}
	return newGrid
}

// String returns a string representation of the grid
func (g *Grid) String() string {
	result := ""
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.IsAlive(x, y) {
				result += "█"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}
