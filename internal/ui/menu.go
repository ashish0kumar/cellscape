package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// MenuModel represents the main menu state
type MenuModel struct {
	choices  []string
	cursor   int
	selected string
	width    int
	height   int
}

// NewMenuModel creates a new MenuModel with default choices and dimensions
func NewMenuModel() MenuModel {
	return MenuModel{
		choices: []string{
			"Conway's Game of Life",
			"Brian's Brain",
			"Langton's Ant",
			"Larger than Life",
			"Belousov-Zhabotinsky",
			"Faders",
			"Forest Fire",
			"Wildfire",
		},
		width:  80,
		height: 24,
	}
}

// SetDimensions allows setting the menu dimensions when returning from a game
func (m *MenuModel) SetDimensions(width, height int) {
	m.width = width
	m.height = height
}

func (m MenuModel) Init() tea.Cmd {
	return tea.WindowSize() // Request window size on init
}

// Update handles user input and updates the menu state
func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	// Handle key messages for navigation and selection
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = m.choices[m.cursor]

			var automatonType string

			// Launch and determine the automaton type based on selection
			switch m.selected {
			case "Conway's Game of Life":
				automatonType = "life"
			case "Langton's Ant":
				automatonType = "ant"
			case "Brian's Brain":
				automatonType = "brain"
			case "Forest Fire":
				automatonType = "forest"
			case "Wildfire":
				automatonType = "wildfire"
			case "Belousov-Zhabotinsky":
				automatonType = "belousov"
			case "Larger than Life":
				automatonType = "ltl"
			case "Faders":
				automatonType = "faders"
			}

			// Create and run simulation
			simModel := NewSimulationModel(automatonType)
			return simModel, simModel.Init()
		}
	}
	return m, nil
}

// View renders the menu
func (m MenuModel) View() string {

	title := titleStyle.Width(m.width).Render(asciiTitle)
	var menuItems []string
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			item := selectedMenuItemStyle.Width(30).Align(lipgloss.Left).Render(fmt.Sprintf("%s %s", cursor, choice))
			menuItems = append(menuItems, item)
		} else {
			item := menuItemStyle.Width(30).Align(lipgloss.Left).Render(fmt.Sprintf("%s %s", cursor, choice))
			menuItems = append(menuItems, item)
		}
	}

	menuBlock := strings.Join(menuItems, "\n")
	centeredMenu := lipgloss.Place(m.width, len(menuItems), lipgloss.Center, lipgloss.Top, menuBlock)
	help := mutedHelpStyle.Width(m.width).Render("Use ↑/↓ arrows to navigate • Enter to select • q to quit")

	contentHeight := 8 + len(menuItems) + 6
	topPadding := (m.height - contentHeight) / 2
	if topPadding < 0 {
		topPadding = 0
	}

	content := title + "\n\n"
	content += centeredMenu
	content += "\n\n\n" + help

	for i := 0; i < topPadding; i++ {
		content = "\n" + content
	}

	return content
}
