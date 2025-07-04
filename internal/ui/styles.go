package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const asciiTitle = `
 ██████╗███████╗██╗     ██╗     ███████╗ ██████╗ █████╗ ██████╗ ███████╗
██╔════╝██╔════╝██║     ██║     ██╔════╝██╔════╝██╔══██╗██╔══██╗██╔════╝
██║     █████╗  ██║     ██║     ███████╗██║     ███████║██████╔╝█████╗  
██║     ██╔══╝  ██║     ██║     ╚════██║██║     ██╔══██║██╔═══╝ ██╔══╝  
╚██████╗███████╗███████╗███████╗███████║╚██████╗██║  ██║██║     ███████╗
 ╚═════╝╚══════╝╚══════╝╚══════╝╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝     ╚══════╝
`

var (
	primaryColor    = lipgloss.Color("12")   // Blue
	mutedColor      = lipgloss.Color("8")    // Grey
	backgroundColor = lipgloss.Color("#000") // Black
	textColor       = lipgloss.Color("15")   // White

	MonochromeMode = false

	generationColors = []lipgloss.Color{
		lipgloss.Color("1"),
		lipgloss.Color("2"),
		lipgloss.Color("3"),
		lipgloss.Color("4"),
		lipgloss.Color("5"),
		lipgloss.Color("6"),
		lipgloss.Color("7"),
		lipgloss.Color("8"),
		lipgloss.Color("9"),
		lipgloss.Color("10"),
		lipgloss.Color("11"),
		lipgloss.Color("12"),
		lipgloss.Color("13"),
		lipgloss.Color("14"),
		lipgloss.Color("15"),
	}

	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Align(lipgloss.Center)

	menuItemStyle = lipgloss.NewStyle().
			Foreground(textColor).
			Padding(0, 2).
			Width(30).
			Align(lipgloss.Left)

	selectedMenuItemStyle = lipgloss.NewStyle().
				Foreground(backgroundColor).
				Background(primaryColor).
				Padding(0, 2).
				Bold(true).
				Width(30).
				Align(lipgloss.Left)

	mutedHelpStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Align(lipgloss.Center)

	statusStyle = lipgloss.NewStyle().
			Foreground(textColor).
			Bold(true).
			Align(lipgloss.Center)
)

// GetGenerationColor returns a color based on generation
func GetGenerationColor(generation int) lipgloss.Color {
	if MonochromeMode {
		return textColor // Simple white text
	}
	return generationColors[generation%len(generationColors)]
}

func ToggleMonochromeMode() {
	MonochromeMode = !MonochromeMode
}
