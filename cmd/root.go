// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package cmd

import (
	"fmt"
	"os"

	"github.com/cloudmanic/tradier/client"
	"github.com/cloudmanic/tradier/config"
	"github.com/spf13/cobra"
)

// jsonOutput controls whether output is rendered as JSON instead of tables.
var jsonOutput bool

var rootCmd = &cobra.Command{
	Use:   "tradier",
	Short: "CLI tool for the Tradier brokerage API",
	Long:  "A command-line interface for interacting with the Tradier brokerage API. Supports account management, market data, trading, watchlists, and streaming.",
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Output raw JSON instead of formatted tables")
}

// Execute runs the root command and exits on error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// loadClientFromConfig reads the config file and returns a configured API client.
func loadClientFromConfig() (*client.Client, *config.Config, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load config: %w\nRun 'tradier init' to configure your API key", err)
	}
	c := client.NewClient(cfg.BaseURL(), cfg.APIKey)
	return c, cfg, nil
}

// printResult outputs data as a formatted table (default) or as JSON (with --json flag).
func printResult(data []byte, tableFunc func([]byte)) {
	if jsonOutput {
		pretty, err := client.PrettyJSON(data)
		if err != nil {
			fmt.Println(string(data))
			return
		}
		fmt.Println(pretty)
		return
	}
	tableFunc(data)
}

// requireAccountID returns the account ID from the flag or config, erroring if neither is set.
func requireAccountID(cmd *cobra.Command, cfg *config.Config) (string, error) {
	accountID, _ := cmd.Flags().GetString("account-id")
	if accountID == "" {
		accountID = cfg.AccountID
	}
	if accountID == "" {
		return "", fmt.Errorf("--account-id is required (or set account_id in config via 'tradier init')")
	}
	return accountID, nil
}
