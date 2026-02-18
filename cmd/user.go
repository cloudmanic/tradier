// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package cmd

import (
	"github.com/spf13/cobra"
)

// userCmd is the parent command for user-related subcommands.
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User management commands",
	Long:  "Commands for managing Tradier user information.",
}

// profileCmd retrieves the profile information for the authenticated user.
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Get user profile information",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		data, err := c.GetProfile()
		if err != nil {
			return err
		}
		printResult(data, displayProfile)
		return nil
	},
}

func init() {
	userCmd.AddCommand(profileCmd)
	rootCmd.AddCommand(userCmd)
}
