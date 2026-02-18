// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tradingCmd is the parent command for all trading subcommands.
var tradingCmd = &cobra.Command{
	Use:   "trading",
	Short: "Trading commands",
	Long:  "Commands for placing, modifying, and canceling orders. Supports equity, option, multileg, combo, OTO, OCO, and OTOCO order types.",
}

// placeOrderCmd places a new trading order.
var placeOrderCmd = &cobra.Command{
	Use:   "place",
	Short: "Place a new trading order",
	Long: `Place a trading order. Supports equity, option, multileg, combo, OTO, OCO, and OTOCO orders.

Examples:
  # Equity market buy
  tradier trading place --class equity --symbol AAPL --side buy --quantity 10 --type market --duration day

  # Option limit buy
  tradier trading place --class option --symbol AAPL --option-symbol AAPL220617C00270000 --side buy_to_open --quantity 5 --type limit --duration day --price 3.50

  # Multileg spread
  tradier trading place --class multileg --symbol AAPL --type debit --duration day --price 1.50 \
    --option-symbol-0 AAPL220617C00270000 --side-0 buy_to_open --quantity-0 1 \
    --option-symbol-1 AAPL220617C00280000 --side-1 sell_to_open --quantity-1 1

  # Preview an order (validates without submitting)
  tradier trading place --class equity --symbol AAPL --side buy --quantity 10 --type market --duration day --preview`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, cfg, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		accountID, err := requireAccountID(cmd, cfg)
		if err != nil {
			return err
		}

		// Build the order params from all flags
		params := map[string]string{}
		flagsToParams := map[string]string{
			"class":            "class",
			"symbol":           "symbol",
			"side":             "side",
			"quantity":         "quantity",
			"type":             "type",
			"duration":         "duration",
			"price":            "price",
			"stop":             "stop",
			"tag":              "tag",
			"option-symbol":    "option_symbol",
			"preview":          "preview",
			// Indexed params for multileg/combo/OTO/OCO/OTOCO
			"option-symbol-0":  "option_symbol[0]",
			"side-0":           "side[0]",
			"quantity-0":       "quantity[0]",
			"option-symbol-1":  "option_symbol[1]",
			"side-1":           "side[1]",
			"quantity-1":       "quantity[1]",
			"option-symbol-2":  "option_symbol[2]",
			"side-2":           "side[2]",
			"quantity-2":       "quantity[2]",
			"option-symbol-3":  "option_symbol[3]",
			"side-3":           "side[3]",
			"quantity-3":       "quantity[3]",
			"symbol-0":         "symbol[0]",
			"symbol-1":         "symbol[1]",
			"symbol-2":         "symbol[2]",
			"type-0":           "type[0]",
			"type-1":           "type[1]",
			"type-2":           "type[2]",
			"price-0":          "price[0]",
			"price-1":          "price[1]",
			"price-2":          "price[2]",
			"stop-0":           "stop[0]",
			"stop-1":           "stop[1]",
			"stop-2":           "stop[2]",
			"duration-0":       "duration[0]",
			"duration-1":       "duration[1]",
			"duration-2":       "duration[2]",
		}

		for flag, param := range flagsToParams {
			val, _ := cmd.Flags().GetString(flag)
			if val != "" {
				params[param] = val
			}
		}

		orderClass := params["class"]
		if orderClass == "" {
			return fmt.Errorf("--class is required (equity, option, multileg, combo, oto, oco, otoco)")
		}

		data, err := c.PlaceOrder(accountID, params)
		if err != nil {
			return err
		}
		printResult(data, displayOrderResult)
		return nil
	},
}

// changeOrderCmd modifies an existing order.
var changeOrderCmd = &cobra.Command{
	Use:   "change",
	Short: "Modify an existing order",
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
		params := map[string]string{}
		for _, flag := range []string{"type", "duration", "price", "stop", "tag"} {
			val, _ := cmd.Flags().GetString(flag)
			if val != "" {
				params[flag] = val
			}
		}
		data, err := c.ChangeOrder(accountID, orderID, params)
		if err != nil {
			return err
		}
		printResult(data, displayOrderResult)
		return nil
	},
}

// cancelOrderCmd cancels an existing order.
var cancelOrderCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel an existing order",
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
		data, err := c.CancelOrder(accountID, orderID)
		if err != nil {
			return err
		}
		printResult(data, displayOrderResult)
		return nil
	},
}

func init() {
	// Account ID for all trading commands
	tradingCmds := []*cobra.Command{placeOrderCmd, changeOrderCmd, cancelOrderCmd}
	for _, cmd := range tradingCmds {
		cmd.Flags().String("account-id", "", "Account ID (defaults to config value)")
	}

	// Place order flags - common
	placeOrderCmd.Flags().String("class", "", "Order class: equity, option, multileg, combo, oto, oco, otoco (required)")
	placeOrderCmd.Flags().String("symbol", "", "Symbol")
	placeOrderCmd.Flags().String("side", "", "Side: buy, sell, sell_short, buy_to_cover, buy_to_open, buy_to_close, sell_to_open, sell_to_close")
	placeOrderCmd.Flags().String("quantity", "", "Quantity")
	placeOrderCmd.Flags().String("type", "", "Order type: market, limit, stop, stop_limit, debit, credit, even")
	placeOrderCmd.Flags().String("duration", "", "Duration: day, gtc, pre, post")
	placeOrderCmd.Flags().String("price", "", "Limit price")
	placeOrderCmd.Flags().String("stop", "", "Stop price")
	placeOrderCmd.Flags().String("tag", "", "User-defined order tag")
	placeOrderCmd.Flags().String("option-symbol", "", "OCC option symbol (for single option orders)")
	placeOrderCmd.Flags().String("preview", "", "Preview order without submitting: true/false")

	// Indexed leg params for multileg/combo/OTO/OCO/OTOCO orders
	for i := 0; i <= 3; i++ {
		suffix := fmt.Sprintf("-%d", i)
		placeOrderCmd.Flags().String("option-symbol"+suffix, "", fmt.Sprintf("Leg %d option symbol", i))
		placeOrderCmd.Flags().String("side"+suffix, "", fmt.Sprintf("Leg %d side", i))
		placeOrderCmd.Flags().String("quantity"+suffix, "", fmt.Sprintf("Leg %d quantity", i))
		placeOrderCmd.Flags().String("symbol"+suffix, "", fmt.Sprintf("Leg %d symbol", i))
		placeOrderCmd.Flags().String("type"+suffix, "", fmt.Sprintf("Leg %d order type", i))
		placeOrderCmd.Flags().String("price"+suffix, "", fmt.Sprintf("Leg %d price", i))
		placeOrderCmd.Flags().String("stop"+suffix, "", fmt.Sprintf("Leg %d stop price", i))
		placeOrderCmd.Flags().String("duration"+suffix, "", fmt.Sprintf("Leg %d duration", i))
	}

	// Change order flags
	changeOrderCmd.Flags().String("order-id", "", "Order ID to modify (required)")
	changeOrderCmd.Flags().String("type", "", "New order type: market, limit, stop, stop_limit")
	changeOrderCmd.Flags().String("duration", "", "New duration: day, gtc, pre, post")
	changeOrderCmd.Flags().String("price", "", "New limit price")
	changeOrderCmd.Flags().String("stop", "", "New stop price")
	changeOrderCmd.Flags().String("tag", "", "New user-defined order tag")

	// Cancel order flags
	cancelOrderCmd.Flags().String("order-id", "", "Order ID to cancel (required)")

	// Build command tree
	tradingCmd.AddCommand(placeOrderCmd, changeOrderCmd, cancelOrderCmd)
	rootCmd.AddCommand(tradingCmd)
}
