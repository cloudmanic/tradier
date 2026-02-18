// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// accountsCmd is the parent command for all account-related subcommands.
var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Account management commands",
	Long:  "Commands for managing Tradier brokerage accounts including balances, positions, orders, and history.",
}

// balanceCmd retrieves the current balance and margin information for an account.
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Get account balance and margin information",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		data, err := c.GetBalances(accountID)
		if err != nil {
			return err
		}
		printResult(data, displayBalance)
		return nil
	},
}

// gainlossCmd retrieves cost basis and gain/loss information for closed positions.
var gainlossCmd = &cobra.Command{
	Use:   "gainloss",
	Short: "Get cost basis and gain/loss for closed positions",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		page, _ := cmd.Flags().GetString("page")
		limit, _ := cmd.Flags().GetString("limit")
		sortBy, _ := cmd.Flags().GetString("sort-by")
		sort, _ := cmd.Flags().GetString("sort")
		data, err := c.GetGainLoss(accountID, page, limit, sortBy, sort)
		if err != nil {
			return err
		}
		printResult(data, displayGainLoss)
		return nil
	},
}

// historicalBalancesCmd retrieves historical account balances over time.
var historicalBalancesCmd = &cobra.Command{
	Use:   "historical-balances",
	Short: "Get historical account balances over time",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		period, _ := cmd.Flags().GetString("period")
		data, err := c.GetHistoricalBalances(accountID, period)
		if err != nil {
			return err
		}
		printResult(data, displayHistoricalBalances)
		return nil
	},
}

// historyCmd retrieves historical activity for an account.
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Get historical activity for an account",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		page, _ := cmd.Flags().GetString("page")
		limit, _ := cmd.Flags().GetString("limit")
		activityType, _ := cmd.Flags().GetString("type")
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		data, err := c.GetHistory(accountID, page, limit, activityType, start, end)
		if err != nil {
			return err
		}
		printResult(data, displayHistory)
		return nil
	},
}

// orderCmd retrieves a specific order by its ID.
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Get a specific order by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		orderID, _ := cmd.Flags().GetString("order-id")
		if orderID == "" {
			return fmt.Errorf("--order-id is required")
		}
		includeTags, _ := cmd.Flags().GetString("include-tags")
		data, err := c.GetOrder(accountID, orderID, includeTags)
		if err != nil {
			return err
		}
		printResult(data, displayOrder)
		return nil
	},
}

// ordersCmd retrieves all orders for an account.
var ordersCmd = &cobra.Command{
	Use:   "orders",
	Short: "Get all orders for an account",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		page, _ := cmd.Flags().GetString("page")
		limit, _ := cmd.Flags().GetString("limit")
		includeTags, _ := cmd.Flags().GetString("include-tags")
		data, err := c.GetOrders(accountID, page, limit, includeTags)
		if err != nil {
			return err
		}
		printResult(data, displayOrders)
		return nil
	},
}

// positionsCmd retrieves current positions held in an account.
var positionsCmd = &cobra.Command{
	Use:   "positions",
	Short: "Get current positions in an account",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		data, err := c.GetPositions(accountID)
		if err != nil {
			return err
		}
		printResult(data, displayPositions)
		return nil
	},
}

// positionGroupsCmd is the parent command for position group operations.
var positionGroupsCmd = &cobra.Command{
	Use:   "position-groups",
	Short: "Manage position groups",
	Long:  "Commands for listing, creating, updating, and deleting position groups.",
}

// listPositionGroupsCmd retrieves all position groups for an account.
var listPositionGroupsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all position groups",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		data, err := c.GetPositionGroups(accountID)
		if err != nil {
			return err
		}
		printResult(data, displayPositionGroups)
		return nil
	},
}

// createPositionGroupCmd creates a new position group with the given label and symbols.
var createPositionGroupCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new position group",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		label, _ := cmd.Flags().GetString("label")
		symbols, _ := cmd.Flags().GetString("symbols")
		if label == "" || symbols == "" {
			return fmt.Errorf("--label and --symbols are required")
		}
		data, err := c.CreatePositionGroup(accountID, label, symbols)
		if err != nil {
			return err
		}
		printResult(data, displayPositionGroup)
		return nil
	},
}

// updatePositionGroupCmd updates an existing position group.
var updatePositionGroupCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a position group",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		groupID, _ := cmd.Flags().GetString("group-id")
		label, _ := cmd.Flags().GetString("label")
		symbols, _ := cmd.Flags().GetString("symbols")
		if groupID == "" || label == "" || symbols == "" {
			return fmt.Errorf("--group-id, --label, and --symbols are required")
		}
		data, err := c.UpdatePositionGroup(accountID, groupID, label, symbols)
		if err != nil {
			return err
		}
		printResult(data, displayPositionGroup)
		return nil
	},
}

// deletePositionGroupCmd deletes a position group from an account.
var deletePositionGroupCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a position group",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}
		groupID, _ := cmd.Flags().GetString("group-id")
		if groupID == "" {
			return fmt.Errorf("--group-id is required")
		}
		data, err := c.DeletePositionGroup(accountID, groupID)
		if err != nil {
			return err
		}
		printResult(data, displayGeneric)
		return nil
	},
}

func init() {
	// Add account-id flag to all account commands
	accountCmds := []*cobra.Command{balanceCmd, gainlossCmd, historicalBalancesCmd, historyCmd, orderCmd, ordersCmd, positionsCmd}
	for _, cmd := range accountCmds {
		cmd.Flags().String("account-id", "", "Account ID (defaults to config value)")
	}

	// Position group commands also need account-id
	pgCmds := []*cobra.Command{listPositionGroupsCmd, createPositionGroupCmd, updatePositionGroupCmd, deletePositionGroupCmd}
	for _, cmd := range pgCmds {
		cmd.Flags().String("account-id", "", "Account ID (defaults to config value)")
	}

	// Gainloss-specific flags
	gainlossCmd.Flags().String("page", "", "Page number for pagination")
	gainlossCmd.Flags().String("limit", "", "Number of results to return")
	gainlossCmd.Flags().String("sort-by", "", "Sort by: closedate, opendate, symbol, gainloss")
	gainlossCmd.Flags().String("sort", "", "Sort direction: asc, desc")

	// Historical balances flags
	historicalBalancesCmd.Flags().String("period", "", "Period: WEEK, MONTH, YTD, YEAR, YEAR_3, YEAR_5, ALL")

	// History flags
	historyCmd.Flags().String("page", "", "Page number for pagination")
	historyCmd.Flags().String("limit", "", "Number of events to return")
	historyCmd.Flags().String("type", "", "Event type: trade, option, ach, wire, dividend, fee, tax, journal, check, transfer, adjustment")
	historyCmd.Flags().String("start", "", "Start date (YYYY-MM-DD)")
	historyCmd.Flags().String("end", "", "End date (YYYY-MM-DD)")

	// Order flags
	orderCmd.Flags().String("order-id", "", "Order ID (required)")
	orderCmd.Flags().String("include-tags", "", "Include user-defined tags: true/false")

	// Orders flags
	ordersCmd.Flags().String("page", "", "Page number for pagination")
	ordersCmd.Flags().String("limit", "", "Number of orders to return")
	ordersCmd.Flags().String("include-tags", "", "Include user-defined tags: true/false")

	// Position group specific flags
	createPositionGroupCmd.Flags().String("label", "", "Position group label (required)")
	createPositionGroupCmd.Flags().String("symbols", "", "Comma-separated list of symbols (required)")
	updatePositionGroupCmd.Flags().String("group-id", "", "Position group ID (required)")
	updatePositionGroupCmd.Flags().String("label", "", "Position group label (required)")
	updatePositionGroupCmd.Flags().String("symbols", "", "Comma-separated list of symbols (required)")
	deletePositionGroupCmd.Flags().String("group-id", "", "Position group ID (required)")

	// Build the command tree
	positionGroupsCmd.AddCommand(listPositionGroupsCmd, createPositionGroupCmd, updatePositionGroupCmd, deletePositionGroupCmd)
	accountsCmd.AddCommand(balanceCmd, gainlossCmd, historicalBalancesCmd, historyCmd, orderCmd, ordersCmd, positionsCmd, positionGroupsCmd)
	rootCmd.AddCommand(accountsCmd)
}
