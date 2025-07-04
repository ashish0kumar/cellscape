package ui

import (
	"cellscape/internal/automata"
	"cellscape/internal/grid"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

// SimulationModel represents the state of the cellular automaton simulation
type SimulationModel struct {
	automaton  automata.Automaton
	grid       *grid.Grid
	running    bool
	generation int
	width      int
	height     int
	cellAges   [][]int
	FocusMode  bool
}

// NewSimulationModel creates a new SimulationModel for the specified automaton type
func NewSimulationModel(automatonType string) SimulationModel {
	return SimulationModel{
		automaton:  automata.NewAutomaton(automatonType),
		grid:       grid.NewGrid(1, 1),
		running:    true,
		generation: 0,
		width:      0,
		height:     0,
		cellAges:   [][]int{},
		FocusMode:  false,
	}
}

// Init initializes the simulation model, setting up the grid and starting the tick
func (m SimulationModel) Init() tea.Cmd {
	return tea.Batch(
		tea.WindowSize(),
		m.tick(),
	)
}

// tick creates a command that ticks the simulation every 60 ms
func (m SimulationModel) tick() tea.Cmd {
	return tea.Tick(time.Millisecond*60, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles user input and updates the simulation state
func (m SimulationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var newWidth, newHeight int
		if m.FocusMode {
			// In focus mode, use entire screen
			newWidth = msg.Width
			newHeight = msg.Height
		} else {
			// Normal mode, leave space for status and help
			newWidth = msg.Width
			newHeight = msg.Height - 4
		}

		// Always initialize if not done yet, or resize if different
		if newWidth > 10 && newHeight > 5 && (newWidth != m.width || newHeight != m.height) {
			m.width = newWidth
			m.height = newHeight
			m.grid = grid.NewGrid(newWidth, newHeight)
			m.automaton.Initialize(m.grid)
			m.generation = 0

			// Initialize cell ages
			m.cellAges = make([][]int, newHeight)
			for i := range m.cellAges {
				m.cellAges[i] = make([]int, newWidth)
			}
		}

	// Handle key messages for user input
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q", "esc":
			// Return to menu with proper window size
			menuModel := NewMenuModel()
			menuModel.width = m.width
			menuModel.height = m.height
			return menuModel, tea.WindowSize()

		case " ":
			// Pause resume
			m.running = !m.running

		case "f":
			// Toggle focus mode
			m.FocusMode = !m.FocusMode

			// Force resize to adjust grid size
			return m, tea.WindowSize()

		case "c":
			// Toggle colorful mode
			ToggleMonochromeMode()

		case "r":
			// Reset
			m.grid.Clear()
			m.automaton.Initialize(m.grid)
			m.generation = 0

			// Reset cell ages
			for y := 0; y < len(m.cellAges); y++ {
				for x := 0; x < len(m.cellAges[y]); x++ {
					m.cellAges[y][x] = 0
				}
			}

		case "s":
			// Step only if running is false
			if !m.running {
				m.stepSimulation()
			}
		}

	// Handle tick messages to update the simulation
	case tickMsg:
		if m.running && len(m.cellAges) > 0 {
			m.stepSimulation()
		}
		return m, m.tick()
	}

	return m, nil
}

// stepSimulation performs a single step of the simulation, updating the grid and cell ages
func (m *SimulationModel) stepSimulation() {
	if m.grid == nil || len(m.cellAges) == 0 {
		return
	}

	oldGrid := m.grid.Copy()
	m.grid = m.automaton.Step(m.grid)
	m.generation++

	// Update cell ages for color cycling
	for y := 0; y < m.grid.Height; y++ {
		for x := 0; x < m.grid.Width; x++ {

			// Special case for B-Z automaton
			if m.automaton.Name() == "Belousov-Zhabotinsky" {
				currentValue := m.grid.Get(x, y).Value
				oldValue := oldGrid.Get(x, y).Value

				if currentValue > 0 {
					if oldValue == 0 {
						// New reaction started
						m.cellAges[y][x] = m.generation
					}
					// Keep existing age for continuing reactions
				} else {
					// No reaction, reset age
					m.cellAges[y][x] = 0
				}
			} else {
				// for other automata
				if m.grid.IsAlive(x, y) {
					if !oldGrid.IsAlive(x, y) {
						// Assign birth generation if cell was born this step
						m.cellAges[y][x] = m.generation
					}
					// Existing cells keep their birth generation
				} else {
					// Cell is dead, reset age
					m.cellAges[y][x] = 0
				}
			}
		}
	}
}

// View renders the current state of the simulation
func (m SimulationModel) View() string {
	if len(m.cellAges) == 0 {
		return "Loading..."
	}

	// Grid visualization
	var gridBuilder strings.Builder
	for y := 0; y < m.grid.Height; y++ {
		for x := 0; x < m.grid.Width; x++ {
			cell := m.grid.Get(x, y)
			if cell.Alive || cell.Value > 0 {
				char, color := m.getCellDisplay(cell, x, y)
				coloredCell := lipgloss.NewStyle().Foreground(color).Render(char)
				gridBuilder.WriteString(coloredCell)
			} else {
				gridBuilder.WriteString(" ")
			}
		}
		if y < m.grid.Height-1 {
			gridBuilder.WriteString("\n")
		}
	}

	gridDisplay := gridBuilder.String()

	// In focus mode, return only the grid
	if m.FocusMode {
		return gridDisplay
	}

	// set mode text based on MonochromeMode
	modeText := "Colorful"
	if MonochromeMode {
		modeText = "Monochrome"
	}

	// In normal mode, add status and help text

	// status text
	status := fmt.Sprintf("%s | Generation: %d | %s | Mode: %s",
		m.automaton.Name(),
		m.generation,
		func() string {
			if m.running {
				return "Running"
			}
			return "Paused"
		}(),
		modeText)

	statusBar := statusStyle.Width(m.width).Render(status)

	// Help text
	help := mutedHelpStyle.Width(m.width).Render("Space: pause/resume • S: step • R: reset • F: focus • C: colors • Q/Esc: back")

	content := statusBar + "\n\n" + gridDisplay + "\n\n" + help
	return content
}

// getCellDisplay returns the character and color for a cell based on its state
func (m SimulationModel) getCellDisplay(cell grid.Cell, x, y int) (string, lipgloss.Color) {
	switch m.automaton.Name() {

	// Special cases for specific automata
	case "Brian's Brain":
		if MonochromeMode {
			switch cell.Value {
			case 1:
				return "█", lipgloss.Color("15") // Live
			case 2:
				return "▓", lipgloss.Color("8") // Dying
			default:
				return "░", lipgloss.Color("0") // Dead
			}
		} else {
			switch cell.Value {
			case 1:
				return "█", lipgloss.Color("2") // Live
			case 2:
				return "▓", lipgloss.Color("8") // Dying
			default:
				return "░", lipgloss.Color("0") // Dead
			}
		}
	case "Forest Fire":
		if MonochromeMode {
			switch cell.Value {
			case 1:
				return "█", lipgloss.Color("8") // Tree
			case 2:
				return "█", lipgloss.Color("15") // Fire
			default:
				return " ", lipgloss.Color("0") // Empty
			}
		} else {
			switch cell.Value {
			case 1:
				return "█", lipgloss.Color("#228B22") // Tree
			case 2:
				return "█", lipgloss.Color("#FF4500") // Fire
			default:
				return " ", lipgloss.Color("0") // Empty
			}
		}
	case "Wildfire":
		if MonochromeMode {
			switch cell.Value {
			case 1:
				return "█", lipgloss.Color("8") // Vegetation
			case 2:
				return "█", lipgloss.Color("15") // Fire
			case 3:
				return " ", lipgloss.Color("0") // Burnt
			default:
				return " ", lipgloss.Color("0") // Empty
			}
		} else {
			switch cell.Value {
			case 1:
				return "█", lipgloss.Color("#228B22") // Vegetation
			case 2:
				return "█", lipgloss.Color("#FF4500") // Fire
			case 3:
				return " ", lipgloss.Color("0") // Burnt
			default:
				return " ", lipgloss.Color("0") // Empty - black
			}
		}
	case "Belousov-Zhabotinsky":
		if MonochromeMode {
			intensity := cell.Value
			if intensity > 255 {
				intensity = 255
			}
			// Map intensity to grayscale ANSI
			grayLevel := 232 + (intensity * 23 / 255)
			return "█", lipgloss.Color(fmt.Sprintf("%d", grayLevel))
		} else {
			intensity := cell.Value
			if intensity > 255 {
				intensity = 255
			}
			color := lipgloss.Color(fmt.Sprintf("#%02X%02X%02X", intensity, intensity/2, 255-intensity))
			return "█", color
		}

	// Default case for other automata
	default:
		// Use generation based coloring
		if cell.Alive {
			cellAge := m.cellAges[y][x]
			color := GetGenerationColor(cellAge)
			return "█", color
		}
		return " ", lipgloss.Color("0")
	}
}
