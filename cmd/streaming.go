// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package cmd

import (
	"github.com/spf13/cobra"
)

// streamingCmd is the parent command for streaming session subcommands.
var streamingCmd = &cobra.Command{
	Use:   "streaming",
	Short: "Streaming session commands",
	Long:  "Commands for creating streaming sessions for real-time market data and account events.",
}

// marketSessionCmd creates a streaming session for real-time market data.
var marketSessionCmd = &cobra.Command{
	Use:   "market-session",
	Short: "Create a market data streaming session",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		data, err := c.CreateMarketSession()
		if err != nil {
			return err
		}
		printResult(data, displaySession)
		return nil
	},
}

// accountSessionCmd creates a streaming session for real-time account events.
var accountSessionCmd = &cobra.Command{
	Use:   "account-session",
	Short: "Create an account events streaming session",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		data, err := c.CreateAccountSession()
		if err != nil {
			return err
		}
		printResult(data, displaySession)
		return nil
	},
}

func init() {
	streamingCmd.AddCommand(marketSessionCmd, accountSessionCmd)
	rootCmd.AddCommand(streamingCmd)
}
