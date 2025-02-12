package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pedro-coelho-dr/mdout/config"
	"github.com/spf13/cobra"
)

var (
	captureFlag  string
	outputFlag   string
	languageFlag string
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View or update the current mdout configuration",
	Long: `Displays the current mdout configuration and lets you update settings
using flags. You can set the command capture method (default "none"),
change the output file (if no ".md" extension is provided, ".md" will be appended),
and set the snippet language for Markdown output (e.g., bash, go, etc.).`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()

		// capture flag
		if captureFlag != "" {
			// Accept only "none" or "zsh" for now.
			if captureFlag == "none" || captureFlag == "zsh" {
				cfg.CommandCapture = captureFlag
			} else {
				fmt.Println("Invalid capture method. Valid options are: none, zsh")
				return
			}
		}
		//output flag
		if outputFlag != "" {
			// Append ".md" if no extension is provided.
			if filepath.Ext(outputFlag) != ".md" {
				outputFlag += ".md"
			}
			cfg.OutputFile = outputFlag
		}
		//language flag
		if languageFlag != "" {
			cfg.SnippetLanguage = languageFlag
		}

		//save
		config.SaveConfig(cfg)

		// Print
		fmt.Println("Current mdout Configuration:")
		fmt.Printf("  Command Capture: %s\n", cfg.CommandCapture)
		fmt.Printf("  Snippet Language: ```%s```\n", cfg.SnippetLanguage)
		fmt.Printf("  Output File: %s\n", cfg.OutputFile)
	},
}

func init() {
	configCmd.Flags().StringVarP(&captureFlag, "capture", "c", "", "Set command capture method (none, zsh)")
	configCmd.Flags().StringVarP(&outputFlag, "output", "o", "", "Set output file name (e.g., mdout.md or mdout)")
	configCmd.Flags().StringVarP(&languageFlag, "language", "l", "", "Set snippet language for Markdown output (e.g., bash, console, go, etc.)")

	rootCmd.AddCommand(configCmd)
}
