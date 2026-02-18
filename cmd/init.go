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

// initCmd initializes the tradier configuration by prompting for an API key and storing it in ~/.config/tradier/.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the Tradier CLI configuration",
	Long:  "Prompts for your Tradier API key and optional account ID, then saves the configuration to ~/.config/tradier/config.json.",
	RunE:  runInit,
}

func init() {
	initCmd.Flags().Bool("sandbox", false, "Use the Tradier sandbox environment instead of production")
	rootCmd.AddCommand(initCmd)
}

// runInit prompts the user for API credentials and saves them to the config file.
func runInit(cmd *cobra.Command, args []string) error {
	sandbox, _ := cmd.Flags().GetBool("sandbox")

	reader := bufio.NewReader(os.Stdin)

	if sandbox {
		fmt.Println("Configuring for Tradier SANDBOX environment")
	} else {
		fmt.Println("Configuring for Tradier PRODUCTION environment")
	}

	fmt.Print("Enter your Tradier API key: ")
	apiKey, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read API key: %w", err)
	}
	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return fmt.Errorf("API key cannot be empty")
	}

	fmt.Print("Enter your default account ID (optional, press Enter to skip): ")
	accountID, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read account ID: %w", err)
	}
	accountID = strings.TrimSpace(accountID)

	cfg := &config.Config{
		APIKey:    apiKey,
		Sandbox:   sandbox,
		AccountID: accountID,
	}

	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	configPath, _ := config.ConfigFilePath()
	fmt.Printf("Configuration saved to %s\n", configPath)
	if sandbox {
		fmt.Printf("Base URL: %s\n", config.SandboxBaseURL)
	} else {
		fmt.Printf("Base URL: %s\n", config.ProductionBaseURL)
	}

	return nil
}
