// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// marketsCmd is the parent command for all market data subcommands.
var marketsCmd = &cobra.Command{
	Use:   "markets",
	Short: "Market data commands",
	Long:  "Commands for retrieving market data including quotes, options, history, and calendar information.",
}

// quotesCmd retrieves real-time quotes for one or more symbols.
var quotesCmd = &cobra.Command{
	Use:   "quotes",
	Short: "Get quotes for one or more symbols",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		symbols, _ := cmd.Flags().GetString("symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required")
		}
		greeks, _ := cmd.Flags().GetString("greeks")
		data, err := c.GetQuotes(symbols, greeks)
		if err != nil {
			return err
		}
		printResult(data, displayQuotes)
		return nil
	},
}

// postQuotesCmd retrieves quotes for a larger list of symbols via POST.
var postQuotesCmd = &cobra.Command{
	Use:   "post-quotes",
	Short: "Get quotes for a large list of symbols (via POST)",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		symbols, _ := cmd.Flags().GetString("symbols")
		if symbols == "" {
			return fmt.Errorf("--symbols is required")
		}
		greeks, _ := cmd.Flags().GetString("greeks")
		data, err := c.PostQuotes(symbols, greeks)
		if err != nil {
			return err
		}
		printResult(data, displayQuotes)
		return nil
	},
}

// optionsChainsCmd retrieves option chains for a specific symbol and expiration.
var optionsChainsCmd = &cobra.Command{
	Use:   "options-chains",
	Short: "Get option chains for a symbol and expiration date",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		symbol, _ := cmd.Flags().GetString("symbol")
		expiration, _ := cmd.Flags().GetString("expiration")
		if symbol == "" || expiration == "" {
			return fmt.Errorf("--symbol and --expiration are required")
		}
		greeks, _ := cmd.Flags().GetString("greeks")
		data, err := c.GetOptionsChains(symbol, expiration, greeks)
		if err != nil {
			return err
		}
		printResult(data, displayOptionsChains)
		return nil
	},
}

// optionsExpirationsCmd retrieves available expiration dates for an underlying symbol.
var optionsExpirationsCmd = &cobra.Command{
	Use:   "options-expirations",
	Short: "Get available expiration dates for a symbol",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		symbol, _ := cmd.Flags().GetString("symbol")
		if symbol == "" {
			return fmt.Errorf("--symbol is required")
		}
		includeAllRoots, _ := cmd.Flags().GetString("include-all-roots")
		strikes, _ := cmd.Flags().GetString("strikes")
		contractSize, _ := cmd.Flags().GetString("contract-size")
		expirationType, _ := cmd.Flags().GetString("expiration-type")
		data, err := c.GetOptionsExpirations(symbol, includeAllRoots, strikes, contractSize, expirationType)
		if err != nil {
			return err
		}
		printResult(data, displayOptionsExpirations)
		return nil
	},
}

// optionsStrikesCmd retrieves available strike prices for a symbol and expiration.
var optionsStrikesCmd = &cobra.Command{
	Use:   "options-strikes",
	Short: "Get available strike prices for a symbol and expiration",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		symbol, _ := cmd.Flags().GetString("symbol")
		expiration, _ := cmd.Flags().GetString("expiration")
		if symbol == "" || expiration == "" {
			return fmt.Errorf("--symbol and --expiration are required")
		}
		data, err := c.GetOptionsStrikes(symbol, expiration)
		if err != nil {
			return err
		}
		printResult(data, displayOptionsStrikes)
		return nil
	},
}

// optionsLookupCmd retrieves all options symbols for a given underlying.
var optionsLookupCmd = &cobra.Command{
	Use:   "options-lookup",
	Short: "Look up all options symbols for an underlying",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		underlying, _ := cmd.Flags().GetString("underlying")
		if underlying == "" {
			return fmt.Errorf("--underlying is required")
		}
		strike, _ := cmd.Flags().GetString("strike")
		expiration, _ := cmd.Flags().GetString("expiration")
		optionType, _ := cmd.Flags().GetString("type")
		data, err := c.GetOptionsLookup(underlying, strike, expiration, optionType)
		if err != nil {
			return err
		}
		printResult(data, displayOptionsLookup)
		return nil
	},
}

// marketHistoryCmd retrieves historical OHLCV pricing for a security.
var marketHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get historical pricing for a security",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		symbol, _ := cmd.Flags().GetString("symbol")
		if symbol == "" {
			return fmt.Errorf("--symbol is required")
		}
		interval, _ := cmd.Flags().GetString("interval")
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		data, err := c.GetHistoricalPricing(symbol, interval, start, end)
		if err != nil {
			return err
		}
		printResult(data, displayMarketHistory)
		return nil
	},
}

// timesalesCmd retrieves time and sales data for charting.
var timesalesCmd = &cobra.Command{
	Use:   "timesales",
	Short: "Get time and sales data for a symbol",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		symbol, _ := cmd.Flags().GetString("symbol")
		if symbol == "" {
			return fmt.Errorf("--symbol is required")
		}
		interval, _ := cmd.Flags().GetString("interval")
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		sessionFilter, _ := cmd.Flags().GetString("session-filter")
		data, err := c.GetTimeSales(symbol, interval, start, end, sessionFilter)
		if err != nil {
			return err
		}
		printResult(data, displayTimeSales)
		return nil
	},
}

// calendarCmd retrieves the market calendar for a specific month/year.
var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Get market calendar for a month",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		month, _ := cmd.Flags().GetString("month")
		year, _ := cmd.Flags().GetString("year")
		data, err := c.GetCalendar(month, year)
		if err != nil {
			return err
		}
		printResult(data, displayCalendar)
		return nil
	},
}

// clockCmd retrieves the current intraday market status.
var clockCmd = &cobra.Command{
	Use:   "clock",
	Short: "Get current market clock status",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		data, err := c.GetClock()
		if err != nil {
			return err
		}
		printResult(data, displayClock)
		return nil
	},
}

// etbCmd retrieves the list of easy-to-borrow securities.
var etbCmd = &cobra.Command{
	Use:   "etb",
	Short: "Get Easy-To-Borrow securities list",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		data, err := c.GetETB()
		if err != nil {
			return err
		}
		printResult(data, displayETB)
		return nil
	},
}

// lookupCmd searches for a symbol using the ticker or partial symbol.
var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Look up a symbol by ticker or partial match",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		query, _ := cmd.Flags().GetString("query")
		if query == "" {
			return fmt.Errorf("--query is required")
		}
		exchanges, _ := cmd.Flags().GetString("exchanges")
		types, _ := cmd.Flags().GetString("types")
		data, err := c.GetLookup(query, exchanges, types)
		if err != nil {
			return err
		}
		printResult(data, displaySecurities)
		return nil
	},
}

// searchCmd searches for securities by partial match on symbol or company name.
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for securities by name or symbol",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, _, err := loadClientFromConfig()
		if err != nil {
			return err
		}
		query, _ := cmd.Flags().GetString("query")
		if query == "" {
			return fmt.Errorf("--query is required")
		}
		indexes, _ := cmd.Flags().GetString("indexes")
		data, err := c.GetSearch(query, indexes)
		if err != nil {
			return err
		}
		printResult(data, displaySecurities)
		return nil
	},
}

func init() {
	// Quotes flags
	quotesCmd.Flags().String("symbols", "", "Comma-separated list of symbols (required)")
	quotesCmd.Flags().String("greeks", "", "Include greeks: true/false")

	// Post quotes flags
	postQuotesCmd.Flags().String("symbols", "", "Comma-separated list of symbols (required)")
	postQuotesCmd.Flags().String("greeks", "", "Include greeks: true/false")

	// Options chains flags
	optionsChainsCmd.Flags().String("symbol", "", "Underlying symbol (required)")
	optionsChainsCmd.Flags().String("expiration", "", "Expiration date YYYY-MM-DD (required)")
	optionsChainsCmd.Flags().String("greeks", "", "Include greeks: true/false")

	// Options expirations flags
	optionsExpirationsCmd.Flags().String("symbol", "", "Underlying symbol (required)")
	optionsExpirationsCmd.Flags().String("include-all-roots", "", "Include all option roots: true/false")
	optionsExpirationsCmd.Flags().String("strikes", "", "Include strikes: true/false")
	optionsExpirationsCmd.Flags().String("contract-size", "", "Include contract size: true/false")
	optionsExpirationsCmd.Flags().String("expiration-type", "", "Include expiration type: true/false")

	// Options strikes flags
	optionsStrikesCmd.Flags().String("symbol", "", "Underlying symbol (required)")
	optionsStrikesCmd.Flags().String("expiration", "", "Expiration date YYYY-MM-DD (required)")

	// Options lookup flags
	optionsLookupCmd.Flags().String("underlying", "", "Underlying symbol (required)")
	optionsLookupCmd.Flags().String("strike", "", "Strike price filter")
	optionsLookupCmd.Flags().String("expiration", "", "Expiration date filter YYYY-MM-DD")
	optionsLookupCmd.Flags().String("type", "", "Option type: call, put")

	// History flags
	marketHistoryCmd.Flags().String("symbol", "", "Security symbol (required)")
	marketHistoryCmd.Flags().String("interval", "", "Interval: daily, weekly, monthly")
	marketHistoryCmd.Flags().String("start", "", "Start date YYYY-MM-DD")
	marketHistoryCmd.Flags().String("end", "", "End date YYYY-MM-DD")

	// Timesales flags
	timesalesCmd.Flags().String("symbol", "", "Security symbol (required)")
	timesalesCmd.Flags().String("interval", "", "Interval: tick, 1min, 5min, 15min")
	timesalesCmd.Flags().String("start", "", "Start datetime YYYY-MM-DD HH:MM")
	timesalesCmd.Flags().String("end", "", "End datetime YYYY-MM-DD HH:MM")
	timesalesCmd.Flags().String("session-filter", "", "Session filter: open, all")

	// Calendar flags
	calendarCmd.Flags().String("month", "", "Month (1-12)")
	calendarCmd.Flags().String("year", "", "Year (2000-2050)")

	// Lookup flags
	lookupCmd.Flags().String("query", "", "Search query (required)")
	lookupCmd.Flags().String("exchanges", "", "Exchange codes: Q, N, A, B, C, P, I, M, W, Z")
	lookupCmd.Flags().String("types", "", "Security types: stock, etf, index")

	// Search flags
	searchCmd.Flags().String("query", "", "Search query (required)")
	searchCmd.Flags().String("indexes", "", "Include indexes: true/false")

	// Build the command tree
	marketsCmd.AddCommand(
		quotesCmd, postQuotesCmd,
		optionsChainsCmd, optionsExpirationsCmd, optionsStrikesCmd, optionsLookupCmd,
		marketHistoryCmd, timesalesCmd,
		calendarCmd, clockCmd, etbCmd,
		lookupCmd, searchCmd,
	)
	rootCmd.AddCommand(marketsCmd)
}
