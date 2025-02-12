package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config structure
type Config struct {
	CommandCapture  string `yaml:"command_capture"`  // none, zsh
	SnippetLanguage string `yaml:"snippet_language"` // bash, console, go, etc.
	OutputFile      string `yaml:"output_file"`      // File to write output
}

// Default
var defaultConfig = Config{
	CommandCapture:  "none",
	SnippetLanguage: "bash",
	OutputFile:      "mdout.md",
}

// Get config file path
func getConfigFilePath() string {
	configDir := filepath.Join(os.Getenv("HOME"), ".config")
	return filepath.Join(configDir, "mdout.yaml")
}

// LoadConfig reads the config file, or creates a default one
func LoadConfig() Config {
	configPath := getConfigFilePath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Creating default config at:", configPath)
		SaveConfig(defaultConfig)
		return defaultConfig
	}

	// Read config file
	file, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config:", err)
		return defaultConfig
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		fmt.Println("Error parsing config:", err)
		return defaultConfig
	}

	return cfg
}

// SaveConfig writes the config file
func SaveConfig(cfg Config) {
	configPath := getConfigFilePath()

	// Ensure config directory exists
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Println("Error creating config directory:", err)
		return
	}

	// Marshal and save config
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		fmt.Println("Error encoding config:", err)
		return
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		fmt.Println("Error writing config file:", err)
	}
}
