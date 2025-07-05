package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ashish0kumar/cellscape/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [automaton]",
	Short: "Run a specific cellular automaton",
	Long: `Run a specific cellular automaton directly without going through the menu

Available automata:
  life     - Conway's Game of Life
  ant      - Langton's Ant
  brain    - Brian's Brain
  forest   - Forest Fire
  wildfire - Wildfire
  belousov - Belousov-Zhabotinsky
  ltl      - Larger than Life
  faders   - Faders`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		automatonType := args[0]
		runAutomaton(automatonType)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Add flags for monochrome and focus mode
	runCmd.Flags().BoolVarP(&monochromeFlag, "monochrome", "m", false, "Run in monochrome mode")
	runCmd.Flags().BoolVarP(&focusFlag, "focus", "f", false, "Run in focus mode (full screen, no UI text)")
}

// runAutomaton starts the specified cellular automaton simulation
func runAutomaton(automatonType string) {

	// Define valid automaton types
	var validTypes = map[string]bool{
		"life":     true,
		"ant":      true,
		"brain":    true,
		"forest":   true,
		"wildfire": true,
		"belousov": true,
		"ltl":      true,
		"faders":   true,
	}

	// Check if the provided automaton type is valid
	if !validTypes[automatonType] {
		fmt.Printf("Unknown automaton type: %s\n", automatonType)
		fmt.Println("Available types: life, ant, brain, forest, wildfire, belousov, ltl, faders")
		os.Exit(1)
	}

	// Set modes if flags are provided
	if monochromeFlag {
		ui.ToggleMonochromeMode()
	}

	model := ui.NewSimulationModel(automatonType)

	// Set focus mode if flag is provided
	if focusFlag {
		model.FocusMode = true
	}

	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// startMenu launches the menu for selecting automata
func startMenu() {

	model := ui.NewMenuModel()
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
