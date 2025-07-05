package cmd

import (
	"github.com/ashish0kumar/cellscape/internal/ui"

	"github.com/spf13/cobra"
)

var version = "dev"

var (
	monochromeFlag bool
	focusFlag      bool
)

var rootCmd = &cobra.Command{
	Use:   "cellscape",
	Short: "A terminal-based cellular automata playground",
	Long: `Cellscape is a terminal application for exploring various cellular automata
Navigate through a beautiful ASCII menu and watch automata come to life in your terminal`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {

		// Set modes if flags are provided
		if monochromeFlag {
			ui.ToggleMonochromeMode()
		}

		// If no subcommand is provided, show the menu
		showMenu()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	// Disable the default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Add subcommands
	rootCmd.Flags().BoolVarP(&monochromeFlag, "monochrome", "m", false, "Start in monochrome mode")
	rootCmd.Flags().BoolVarP(&focusFlag, "focus", "f", false, "Start in focus mode (fullscreen, no UI)")
}

func showMenu() {
	startMenu()
}
