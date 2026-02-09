package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-highlights",
	Short: "Generate weekly engineering hightlights from Git/GitHub",
	Long:  `git-highlights analyzes merged PRs and generates meeting-ready markdown summaries of the week's work.`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
