package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pedro-coelho-dr/mdout/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdout",
	Short: "A simple CLI tool to capture piped output and save it into a Markdown file",
	Long:  `mdout is a lightweight command-line tool that allows users to capture command output and save it into a Markdown file.`,
	Run: func(cmd *cobra.Command, args []string) {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			saveFromPipe()
		} else {
			fmt.Println("Usage: command | mdout")
		}
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Initializes Cobra commands and configuration loading
func init() {
}

// Retrieves the last executed command based on the configured shell
func getLastCommand(cfg config.Config) string {
	if cfg.CommandCapture == "none" {
		return ""
	}

	switch cfg.CommandCapture {
	case "zsh":
		return getLastCommandZsh()
	default:
		return "Unknown command"
	}
}

func getLastCommandZsh() string {
	out, err := exec.Command("zsh", "-c", "tail -n 50 ~/.zsh_history").Output()
	if err != nil || len(out) == 0 {
		return "Unknown command"
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	reTimestamp := regexp.MustCompile(`^: ?[0-9]+:[0-9]+;`)
	reBeforeSemicolon := regexp.MustCompile(`^.*?;`)

	var blocks []string
	var currentBlock []string

	// Iterate through the lines in order: oldest --> newest
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if reTimestamp.MatchString(line) {
			if len(currentBlock) > 0 {
				blocks = append(blocks, strings.Join(currentBlock, "\n"))
			}
			cleaned := reBeforeSemicolon.ReplaceAllString(line, "")
			currentBlock = []string{cleaned}
		} else {
			if len(currentBlock) > 0 {
				currentBlock = append(currentBlock, line)
			}
		}
	}
	if len(currentBlock) > 0 {
		blocks = append(blocks, strings.Join(currentBlock, "\n"))
	}
	if len(blocks) == 0 {
		return "Unknown command"
	}

	return blocks[len(blocks)-1]
}

// Saves piped output to a Markdown file and prints it once to the terminal
func saveFromPipe() {
	cfg := config.LoadConfig()

	outputFile := cfg.OutputFile
	if outputFile == "" {
		outputFile = "mdout.md" // Default filename
	}

	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Get the last command if enabled
	lastCommand := getLastCommand(cfg)

	// Capture piped input
	scanner := bufio.NewScanner(os.Stdin)
	var contentBuilder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		contentBuilder.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return
	}

	// Format the snippet for the Markdown file
	output := fmt.Sprintf("\n```%s\n%s\n\n%s```\n",
		cfg.SnippetLanguage, lastCommand, contentBuilder.String())

	// Write the snippet to the file
	if _, err := file.WriteString(output); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
	}
}
