// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// watchlistsCmd is the parent command for all watchlist subcommands.
var watchlistsCmd = &cobra.Command{
	Use:   "watchlists",
	Short: "Watchlist management commands",
	Long:  "Commands for creating, listing, updating, and deleting watchlists and their symbols.",
}

// listWatchlistsCmd retrieves all watchlists for the authenticated user.
var listWatchlistsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all watchlists",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		data, err := c.GetWatchlists()
		if err != nil {
			return err
		}
		printResult(data, displayWatchlists)
		return nil
	},
}

// getWatchlistCmd retrieves a specific watchlist by ID.
var getWatchlistCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a specific watchlist by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			return fmt.Errorf("--id is required")
		}
		data, err := c.GetWatchlist(id)
		if err != nil {
			return err
		}
		printResult(data, displayWatchlist)
		return nil
	},
}

// createWatchlistCmd creates a new watchlist with the given name and symbols.
var createWatchlistCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new watchlist",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		name, _ := cmd.Flags().GetString("name")
		symbols, _ := cmd.Flags().GetString("symbols")
		if name == "" || symbols == "" {
			return fmt.Errorf("--name and --symbols are required")
		}
		data, err := c.CreateWatchlist(name, symbols)
		if err != nil {
			return err
		}
		printResult(data, displayWatchlist)
		return nil
	},
}

// updateWatchlistCmd updates an existing watchlist with a new name and symbols.
var updateWatchlistCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing watchlist",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		symbols, _ := cmd.Flags().GetString("symbols")
		if id == "" || name == "" {
			return fmt.Errorf("--id and --name are required")
		}
		data, err := c.UpdateWatchlist(id, name, symbols)
		if err != nil {
			return err
		}
		printResult(data, displayWatchlist)
		return nil
	},
}

// deleteWatchlistCmd deletes a specific watchlist by ID.
var deleteWatchlistCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a watchlist",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			return fmt.Errorf("--id is required")
		}
		data, err := c.DeleteWatchlist(id)
		if err != nil {
			return err
		}
		printResult(data, displayGeneric)
		return nil
	},
}

// addSymbolsCmd adds symbols to an existing watchlist.
var addSymbolsCmd = &cobra.Command{
	Use:   "add-symbols",
	Short: "Add symbols to a watchlist",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("id")
		symbols, _ := cmd.Flags().GetString("symbols")
		if id == "" || symbols == "" {
			return fmt.Errorf("--id and --symbols are required")
		}
		data, err := c.AddSymbolsToWatchlist(id, symbols)
		if err != nil {
			return err
		}
		printResult(data, displayWatchlist)
		return nil
	},
}

// removeSymbolCmd removes a single symbol from a watchlist.
var removeSymbolCmd = &cobra.Command{
	Use:   "remove-symbol",
	Short: "Remove a symbol from a watchlist",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		id, _ := cmd.Flags().GetString("id")
		symbol, _ := cmd.Flags().GetString("symbol")
		if id == "" || symbol == "" {
			return fmt.Errorf("--id and --symbol are required")
		}
		data, err := c.RemoveSymbolFromWatchlist(id, symbol)
		if err != nil {
			return err
		}
		printResult(data, displayWatchlist)
		return nil
	},
}

func init() {
	// Get watchlist flags
	getWatchlistCmd.Flags().String("id", "", "Watchlist ID (required)")

	// Create watchlist flags
	createWatchlistCmd.Flags().String("name", "", "Watchlist name (required)")
	createWatchlistCmd.Flags().String("symbols", "", "Comma-separated symbols (required)")

	// Update watchlist flags
	updateWatchlistCmd.Flags().String("id", "", "Watchlist ID (required)")
	updateWatchlistCmd.Flags().String("name", "", "Watchlist name (required)")
	updateWatchlistCmd.Flags().String("symbols", "", "Comma-separated symbols")

	// Delete watchlist flags
	deleteWatchlistCmd.Flags().String("id", "", "Watchlist ID (required)")

	// Add symbols flags
	addSymbolsCmd.Flags().String("id", "", "Watchlist ID (required)")
	addSymbolsCmd.Flags().String("symbols", "", "Comma-separated symbols to add (required)")

	// Remove symbol flags
	removeSymbolCmd.Flags().String("id", "", "Watchlist ID (required)")
	removeSymbolCmd.Flags().String("symbol", "", "Symbol to remove (required)")

	// Build command tree
	watchlistsCmd.AddCommand(listWatchlistsCmd, getWatchlistCmd, createWatchlistCmd, updateWatchlistCmd, deleteWatchlistCmd, addSymbolsCmd, removeSymbolCmd)
	rootCmd.AddCommand(watchlistsCmd)
}
