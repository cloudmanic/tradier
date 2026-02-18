// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cloudmanic/tradier/config"
	"github.com/spf13/cobra"
)

// initCmd initializes the tradier configuration by prompting for API keys and storing them in ~/.config/tradier/.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the Tradier CLI configuration",
	Long:  "Prompts for your Tradier production and sandbox API keys and optional account IDs, then saves the configuration to ~/.config/tradier/config.json.",
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// runInit prompts the user for both production and sandbox API credentials and saves them to the config file.
func runInit(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Tradier CLI Configuration")
	fmt.Println("=========================")
	fmt.Println()

	// Production credentials
	fmt.Println("-- Production Environment --")
	fmt.Print("Enter your production API key (press Enter to skip): ")
	prodAPIKey, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}
	prodAPIKey = strings.TrimSpace(prodAPIKey)

	prodAccountID := ""
	if prodAPIKey != "" {
		fmt.Print("Enter your production account ID (optional, press Enter to skip): ")
		prodAccountID, err = readLine(reader)
		if err != nil {
			return err
		}
	}

	fmt.Println()

	// Sandbox credentials
	fmt.Println("-- Sandbox Environment --")
	fmt.Print("Enter your sandbox API key (press Enter to skip): ")
	sandboxAPIKey, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}
	sandboxAPIKey = strings.TrimSpace(sandboxAPIKey)

	sandboxAccountID := ""
	if sandboxAPIKey != "" {
		fmt.Print("Enter your sandbox account ID (optional, press Enter to skip): ")
		sandboxAccountID, err = readLine(reader)
		if err != nil {
			return err
		}
	}

	if prodAPIKey == "" && sandboxAPIKey == "" {
		return fmt.Errorf("at least one API key (production or sandbox) is required")
	}

	cfg := &config.Config{
		ProductionAPIKey:    prodAPIKey,
		ProductionAccountID: prodAccountID,
		SandboxAPIKey:       sandboxAPIKey,
		SandboxAccountID:    sandboxAccountID,
	}

	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	configPath, _ := config.ConfigFilePath()
	fmt.Println()
	fmt.Printf("Configuration saved to %s\n", configPath)

	if prodAPIKey != "" {
		fmt.Printf("Production: %s (key configured)\n", config.ProductionBaseURL)
	}
	if sandboxAPIKey != "" {
		fmt.Printf("Sandbox:    %s (key configured)\n", config.SandboxBaseURL)
	}

	fmt.Println()
	fmt.Println("Use --sandbox flag on any command to use the sandbox environment.")

	return nil
}

// readLine reads a single line from the reader and trims whitespace.
func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	return strings.TrimSpace(line), nil
}
