// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// ===========================================================================
// JSON Navigation Helpers
// ===========================================================================

// parseJSON unmarshals raw JSON bytes into a generic map. Returns nil on failure.
func parseJSON(data []byte) map[string]interface{} {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil
	}
	return m
}

// nested traverses a map by a sequence of keys, returning the nested map or nil.
func nested(data map[string]interface{}, keys ...string) map[string]interface{} {
	current := data
	for _, key := range keys {
		val, ok := current[key]
		if !ok || val == nil {
			return nil
		}
		m, ok := val.(map[string]interface{})
		if !ok {
			return nil
		}
		current = m
	}
	return current
}

// toSlice converts a JSON value to a slice of maps, handling both single objects and arrays.
func toSlice(val interface{}) []map[string]interface{} {
	if val == nil {
		return nil
	}
	switch v := val.(type) {
	case []interface{}:
		result := make([]map[string]interface{}, 0, len(v))
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				result = append(result, m)
			}
		}
		return result
	case map[string]interface{}:
		return []map[string]interface{}{v}
	}
	return nil
}

// toStringSlice converts a JSON array value to a slice of strings.
func toStringSlice(val interface{}) []string {
	if val == nil {
		return nil
	}
	arr, ok := val.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, 0, len(arr))
	for _, item := range arr {
		switch v := item.(type) {
		case string:
			result = append(result, v)
		case float64:
			if v == float64(int64(v)) {
				result = append(result, fmt.Sprintf("%d", int64(v)))
			} else {
				result = append(result, fmt.Sprintf("%.2f", v))
			}
		default:
			result = append(result, fmt.Sprintf("%v", v))
		}
	}
	return result
}

// str safely extracts a string representation from a map value.
func str(m map[string]interface{}, key string) string {
	if m == nil {
		return ""
	}
	val, ok := m[key]
	if !ok || val == nil {
		return ""
	}
	switch v := val.(type) {
	case string:
		return v
	case float64:
		if v == float64(int64(v)) {
			return fmt.Sprintf("%d", int64(v))
		}
		return fmt.Sprintf("%.2f", v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// num safely extracts a float64 from a map value.
func num(m map[string]interface{}, key string) float64 {
	if m == nil {
		return 0
	}
	val, ok := m[key]
	if !ok || val == nil {
		return 0
	}
	f, ok := val.(float64)
	if !ok {
		return 0
	}
	return f
}

// ===========================================================================
// Formatting Helpers
// ===========================================================================

// money formats a float as a dollar amount.
func money(f float64) string {
	if f < 0 {
		return fmt.Sprintf("-$%.2f", -f)
	}
	return fmt.Sprintf("$%.2f", f)
}

// pct formats a float as a percentage with sign.
func pct(f float64) string {
	if f >= 0 {
		return fmt.Sprintf("+%.2f%%", f)
	}
	return fmt.Sprintf("%.2f%%", f)
}

// formatOptionSymbol converts an OCC option symbol (e.g. UNG260220P00014000)
// into a human-readable string (e.g. UNG 02/20/26 $14.00 Put).
// Returns the original string unchanged if it does not match the OCC format.
func formatOptionSymbol(sym string) string {
	// OCC format: ROOT(1-6 chars) + YYMMDD(6) + C/P(1) + STRIKE(8)
	// The last 15 characters are always date + type + strike
	if len(sym) < 15 {
		return sym
	}

	suffix := sym[len(sym)-15:]
	root := sym[:len(sym)-15]

	// Validate date portion is numeric
	datePart := suffix[0:6]
	for _, c := range datePart {
		if c < '0' || c > '9' {
			return sym
		}
	}

	// Validate option type
	optType := suffix[6]
	if optType != 'C' && optType != 'P' {
		return sym
	}

	// Validate strike portion is numeric
	strikePart := suffix[7:15]
	for _, c := range strikePart {
		if c < '0' || c > '9' {
			return sym
		}
	}

	// Parse date: YYMMDD
	yy := datePart[0:2]
	mm := datePart[2:4]
	dd := datePart[4:6]

	// Parse strike: 8 digits representing price * 1000
	var strikeVal int
	for _, c := range strikePart {
		strikeVal = strikeVal*10 + int(c-'0')
	}
	strike := float64(strikeVal) / 1000.0

	typeName := "Call"
	if optType == 'P' {
		typeName = "Put"
	}

	// Format strike: drop decimals if whole number
	var strikeStr string
	if strike == float64(int(strike)) {
		strikeStr = fmt.Sprintf("$%d", int(strike))
	} else {
		strikeStr = fmt.Sprintf("$%.2f", strike)
	}

	return fmt.Sprintf("%s %s/%s/%s %s %s", root, mm, dd, yy, strikeStr, typeName)
}

// shortDate trims an ISO 8601 datetime string to just the date portion.
func shortDate(s string) string {
	if len(s) >= 10 {
		return s[:10]
	}
	return s
}

// ===========================================================================
// Table Printing
// ===========================================================================

// printTable renders a styled table to stdout using go-pretty.
func printTable(headers []string, rows [][]string) {
	if len(rows) == 0 {
		fmt.Println("No results found.")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	headerRow := make(table.Row, len(headers))
	for i, h := range headers {
		headerRow[i] = h
	}
	t.AppendHeader(headerRow)

	for _, row := range rows {
		tableRow := make(table.Row, len(row))
		for i, cell := range row {
			tableRow[i] = cell
		}
		t.AppendRow(tableRow)
	}

	t.SetStyle(table.StyleRounded)
	t.Style().Format.HeaderAlign = text.AlignLeft
	t.Style().Format.Header = text.FormatDefault
	t.Render()
}

// printKV renders a vertical key-value table to stdout using go-pretty.
func printKV(pairs [][2]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	for _, p := range pairs {
		t.AppendRow(table.Row{p[0], p[1]})
	}

	t.SetStyle(table.StyleRounded)
	t.Style().Options.DrawBorder = true
	t.Style().Options.SeparateRows = false
	t.Style().Options.SeparateColumns = true
	t.Style().Options.SeparateHeader = false
	t.Render()
}

// ===========================================================================
// Account Display Functions
// ===========================================================================

// displayBalance renders account balance information as a key-value display.
// Handles margin, pdt, and cash account types which each have their own buying power sub-object.
func displayBalance(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	b := nested(root, "balances")
	if b == nil {
		fmt.Println("No balance data found.")
		return
	}

	pairs := [][2]string{
		{"Account", str(b, "account_number")},
		{"Type", str(b, "account_type")},
		{"Total Equity", money(num(b, "total_equity"))},
		{"Total Cash", money(num(b, "total_cash"))},
		{"Market Value", money(num(b, "market_value"))},
		{"Open P/L", money(num(b, "open_pl"))},
		{"Close P/L", money(num(b, "close_pl"))},
		{"Stock Long Value", money(num(b, "stock_long_value"))},
		{"Option Long Value", money(num(b, "option_long_value"))},
		{"Option Short Value", money(num(b, "option_short_value"))},
		{"Short Market Value", money(num(b, "short_market_value"))},
		{"Current Requirement", money(num(b, "current_requirement"))},
		{"Uncleared Funds", money(num(b, "uncleared_funds"))},
		{"Pending Cash", money(num(b, "pending_cash"))},
		{"Pending Orders", str(b, "pending_orders_count")},
	}
	printKV(pairs)

	// Buying power is nested under the account type key (margin, pdt, or cash)
	accountType := str(b, "account_type")
	bp := nested(b, accountType)
	if bp == nil {
		// Fallback: try common account type keys
		for _, key := range []string{"pdt", "margin", "cash"} {
			if bp = nested(b, key); bp != nil {
				break
			}
		}
	}

	if bp != nil {
		fmt.Println()
		fmt.Println("Buying Power:")
		bpPairs := [][2]string{
			{"Stock Buying Power", money(num(bp, "stock_buying_power"))},
			{"Option Buying Power", money(num(bp, "option_buying_power"))},
		}
		if v := num(bp, "day_trade_buying_power"); v != 0 {
			bpPairs = append(bpPairs, [2]string{"Day Trade Buying Power", money(v)})
		}
		bpPairs = append(bpPairs,
			[2]string{"Fed Call", money(num(bp, "fed_call"))},
			[2]string{"Maintenance Call", money(num(bp, "maintenance_call"))},
		)
		if v := num(bp, "stock_short_value"); v != 0 {
			bpPairs = append(bpPairs, [2]string{"Stock Short Value", money(v)})
		}
		printKV(bpPairs)
	}
}

// displayGainLoss renders closed position gain/loss data as a table.
func displayGainLoss(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	gl := nested(root, "gainloss")
	if gl == nil {
		fmt.Println("No gain/loss data found.")
		return
	}

	positions := toSlice(gl["closed_position"])
	headers := []string{"SYMBOL", "QTY", "COST", "PROCEEDS", "GAIN/LOSS", "GAIN%", "OPEN DATE", "CLOSE DATE"}
	rows := make([][]string, 0, len(positions))
	for _, p := range positions {
		rows = append(rows, []string{
			str(p, "symbol"),
			str(p, "quantity"),
			money(num(p, "cost")),
			money(num(p, "proceeds")),
			money(num(p, "gain_loss")),
			pct(num(p, "gain_loss_percent")),
			shortDate(str(p, "open_date")),
			shortDate(str(p, "close_date")),
		})
	}
	printTable(headers, rows)
}

// displayHistoricalBalances renders historical balance data as a table.
func displayHistoricalBalances(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	balances := toSlice(root["balances"])
	if balances == nil {
		fmt.Println("No historical balance data found.")
		return
	}

	headers := []string{"DATE", "VALUE"}
	rows := make([][]string, 0, len(balances))
	for _, b := range balances {
		rows = append(rows, []string{
			shortDate(str(b, "date")),
			money(num(b, "value")),
		})
	}
	printTable(headers, rows)
}

// displayHistory renders account history events as a table.
// Each event has a type (trade, ach, fee, etc.) with details nested under that type key.
func displayHistory(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	h := nested(root, "history")
	if h == nil {
		fmt.Println("No history data found.")
		return
	}

	events := toSlice(h["event"])
	headers := []string{"DATE", "TYPE", "SYMBOL", "QTY", "PRICE", "AMOUNT", "DESCRIPTION"}
	rows := make([][]string, 0, len(events))
	for _, e := range events {
		eventType := str(e, "type")
		symbol := ""
		qty := ""
		price := ""
		desc := ""

		// Details are nested under the event type key (trade, ach, fee, etc.)
		if detail, ok := e[eventType]; ok {
			if dm, ok := detail.(map[string]interface{}); ok {
				symbol = str(dm, "symbol")
				if q := num(dm, "quantity"); q != 0 {
					qty = fmt.Sprintf("%.0f", q)
				}
				price = money(num(dm, "price"))
				desc = str(dm, "description")
			}
		}

		// Format option symbols for readability
		if symbol != "" {
			symbol = formatOptionSymbol(symbol)
		}

		rows = append(rows, []string{
			shortDate(str(e, "date")),
			eventType,
			symbol,
			qty,
			price,
			money(num(e, "amount")),
			desc,
		})
	}
	printTable(headers, rows)
}

// displayOrder renders a single order as a key-value display.
func displayOrder(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	o := nested(root, "order")
	if o == nil {
		fmt.Println("No order data found.")
		return
	}

	pairs := [][2]string{
		{"Order ID", str(o, "id")},
		{"Class", str(o, "class")},
		{"Symbol", str(o, "symbol")},
		{"Option Symbol", formatOptionSymbol(str(o, "option_symbol"))},
		{"Side", str(o, "side")},
		{"Quantity", str(o, "quantity")},
		{"Type", str(o, "type")},
		{"Price", fmt.Sprintf("%.2f", num(o, "price"))},
		{"Stop", fmt.Sprintf("%.2f", num(o, "stop_price"))},
		{"Status", str(o, "status")},
		{"Duration", str(o, "duration")},
		{"Avg Fill Price", money(num(o, "avg_fill_price"))},
		{"Exec Quantity", str(o, "exec_quantity")},
		{"Remaining", str(o, "remaining_quantity")},
		{"Created", shortDate(str(o, "create_date"))},
	}
	printKV(pairs)

	// Show legs for multileg/combo/advanced orders
	legs := toSlice(o["leg"])
	if len(legs) > 0 {
		fmt.Println()
		fmt.Println("Legs:")
		legHeaders := []string{"OPTION SYMBOL", "SIDE", "QTY", "TYPE", "PRICE", "STATUS", "AVG FILL"}
		legRows := make([][]string, 0, len(legs))
		for _, leg := range legs {
			legRows = append(legRows, []string{
				formatOptionSymbol(str(leg, "option_symbol")),
				str(leg, "side"),
				str(leg, "quantity"),
				str(leg, "type"),
				fmt.Sprintf("%.2f", num(leg, "price")),
				str(leg, "status"),
				money(num(leg, "avg_fill_price")),
			})
		}
		printTable(legHeaders, legRows)
	}
}

// displayOrders renders a list of orders as a table.
func displayOrders(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	o := nested(root, "orders")
	if o == nil {
		fmt.Println("No orders found.")
		return
	}

	orders := toSlice(o["order"])
	headers := []string{"ID", "CLASS", "SYMBOL", "OPTION SYMBOL", "SIDE", "QTY", "TYPE", "PRICE", "STATUS", "DURATION", "AVG FILL", "CREATED"}
	rows := make([][]string, 0, len(orders))
	showLegend := false
	for _, ord := range orders {
		// Format the price column - show limit/stop price when set
		priceStr := ""
		if p := num(ord, "price"); p != 0 {
			priceStr = fmt.Sprintf("%.2f", p)
		}
		if s := num(ord, "stop_price"); s != 0 {
			if priceStr != "" {
				priceStr += "/" + fmt.Sprintf("%.2f", s)
			} else {
				priceStr = fmt.Sprintf("%.2f", s)
			}
		}

		// Get option symbol from the order or from the first leg
		optSym := formatOptionSymbol(str(ord, "option_symbol"))
		hasMultipleLegs := false
		if optSym == "" {
			if legs := toSlice(ord["leg"]); len(legs) > 0 {
				optSym = formatOptionSymbol(str(legs[0], "option_symbol"))
				if len(legs) > 1 {
					optSym += " *"
					hasMultipleLegs = true
				}
			}
		}
		if hasMultipleLegs {
			showLegend = true
		}

		rows = append(rows, []string{
			str(ord, "id"),
			str(ord, "class"),
			str(ord, "symbol"),
			optSym,
			str(ord, "side"),
			str(ord, "quantity"),
			str(ord, "type"),
			priceStr,
			str(ord, "status"),
			str(ord, "duration"),
			money(num(ord, "avg_fill_price")),
			shortDate(str(ord, "create_date")),
		})
	}
	printTable(headers, rows)

	if showLegend {
		fmt.Println("  * = multileg order. Use 'tradier accounts order --order-id <id>' to view all legs.")
	}
}

// displayPositions renders current account positions as a table.
func displayPositions(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	p := nested(root, "positions")
	if p == nil {
		fmt.Println("No positions found.")
		return
	}

	positions := toSlice(p["position"])
	headers := []string{"SYMBOL", "QTY", "COST BASIS", "DATE ACQUIRED"}
	rows := make([][]string, 0, len(positions))
	for _, pos := range positions {
		rows = append(rows, []string{
			str(pos, "symbol"),
			str(pos, "quantity"),
			money(num(pos, "cost_basis")),
			shortDate(str(pos, "date_acquired")),
		})
	}
	printTable(headers, rows)
}

// displayPositionGroups renders all position groups as a table.
func displayPositionGroups(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	// Try common response formats
	groups := toSlice(root["position_groups"])
	if groups == nil {
		pg := nested(root, "positiongroups")
		if pg != nil {
			groups = toSlice(pg["positiongroup"])
		}
	}

	if groups == nil {
		fmt.Println("No position groups found.")
		return
	}

	headers := []string{"ID", "LABEL"}
	rows := make([][]string, 0, len(groups))
	for _, g := range groups {
		rows = append(rows, []string{
			str(g, "id"),
			str(g, "label"),
		})
	}
	printTable(headers, rows)
}

// displayPositionGroup renders a single position group result.
func displayPositionGroup(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	pg := nested(root, "position_group")
	if pg == nil {
		pg = nested(root, "positiongroup")
	}
	if pg == nil {
		// Might be the root itself
		if str(root, "id") != "" || str(root, "label") != "" {
			pg = root
		}
	}

	if pg == nil {
		fmt.Println("Position group updated.")
		return
	}

	pairs := [][2]string{
		{"ID", str(pg, "id")},
		{"Label", str(pg, "label")},
	}
	printKV(pairs)
}

// ===========================================================================
// Market Display Functions
// ===========================================================================

// displayQuotes renders stock/option quote data as a table.
func displayQuotes(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	q := nested(root, "quotes")
	if q == nil {
		fmt.Println("No quote data found.")
		return
	}

	quotes := toSlice(q["quote"])
	headers := []string{"SYMBOL", "LAST", "CHANGE", "CHG%", "VOLUME", "BID", "ASK", "OPEN", "HIGH", "LOW"}
	rows := make([][]string, 0, len(quotes))
	for _, qt := range quotes {
		rows = append(rows, []string{
			str(qt, "symbol"),
			fmt.Sprintf("%.2f", num(qt, "last")),
			fmt.Sprintf("%+.2f", num(qt, "change")),
			pct(num(qt, "change_percentage")),
			str(qt, "volume"),
			fmt.Sprintf("%.2f", num(qt, "bid")),
			fmt.Sprintf("%.2f", num(qt, "ask")),
			fmt.Sprintf("%.2f", num(qt, "open")),
			fmt.Sprintf("%.2f", num(qt, "high")),
			fmt.Sprintf("%.2f", num(qt, "low")),
		})
	}
	printTable(headers, rows)
}

// displayOptionsChains renders option chain data as a table.
func displayOptionsChains(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	o := nested(root, "options")
	if o == nil {
		fmt.Println("No options chain data found.")
		return
	}

	options := toSlice(o["option"])
	headers := []string{"OPTION", "TYPE", "STRIKE", "LAST", "BID", "ASK", "VOLUME", "OPEN INT"}
	rows := make([][]string, 0, len(options))
	for _, opt := range options {
		rows = append(rows, []string{
			formatOptionSymbol(str(opt, "symbol")),
			str(opt, "option_type"),
			fmt.Sprintf("%.2f", num(opt, "strike")),
			fmt.Sprintf("%.2f", num(opt, "last")),
			fmt.Sprintf("%.2f", num(opt, "bid")),
			fmt.Sprintf("%.2f", num(opt, "ask")),
			str(opt, "volume"),
			str(opt, "open_interest"),
		})
	}
	printTable(headers, rows)
}

// displayOptionsExpirations renders available expiration dates as a simple list.
func displayOptionsExpirations(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	exp := nested(root, "expirations")
	if exp == nil {
		fmt.Println("No expiration data found.")
		return
	}

	dates := toStringSlice(exp["date"])
	if len(dates) == 0 {
		// Try alternate format where expirations contain objects
		exps := toSlice(exp["expiration"])
		for _, e := range exps {
			dates = append(dates, shortDate(str(e, "date")))
		}
	}

	headers := []string{"EXPIRATION"}
	rows := make([][]string, 0, len(dates))
	for _, d := range dates {
		rows = append(rows, []string{d})
	}
	printTable(headers, rows)
}

// displayOptionsStrikes renders available strike prices as a simple list.
func displayOptionsStrikes(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	s := nested(root, "strikes")
	if s == nil {
		fmt.Println("No strikes data found.")
		return
	}

	strikes := toStringSlice(s["strike"])
	headers := []string{"STRIKE"}
	rows := make([][]string, 0, len(strikes))
	for _, st := range strikes {
		rows = append(rows, []string{st})
	}
	printTable(headers, rows)
}

// displayOptionsLookup renders options symbol lookup results as a table.
func displayOptionsLookup(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	// Try to find the symbols array in common response structures
	var symbols []map[string]interface{}
	if s := toSlice(root["symbols"]); s != nil {
		symbols = s
	} else if s := nested(root, "symbols"); s != nil {
		symbols = toSlice(s["option"])
	}

	if symbols == nil {
		fmt.Println("No options symbols found.")
		return
	}

	headers := []string{"OPTION", "ROOT", "STRIKE", "EXPIRATION", "TYPE"}
	rows := make([][]string, 0, len(symbols))
	for _, sym := range symbols {
		rows = append(rows, []string{
			formatOptionSymbol(str(sym, "symbol")),
			str(sym, "rootSymbol"),
			fmt.Sprintf("%.2f", num(sym, "strike")),
			shortDate(str(sym, "expiration_date")),
			str(sym, "option_type"),
		})
	}
	printTable(headers, rows)
}

// displayMarketHistory renders historical OHLCV pricing data as a table.
func displayMarketHistory(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	h := nested(root, "history")
	if h == nil {
		fmt.Println("No historical data found.")
		return
	}

	days := toSlice(h["day"])
	headers := []string{"DATE", "OPEN", "HIGH", "LOW", "CLOSE", "VOLUME"}
	rows := make([][]string, 0, len(days))
	for _, d := range days {
		rows = append(rows, []string{
			shortDate(str(d, "date")),
			fmt.Sprintf("%.2f", num(d, "open")),
			fmt.Sprintf("%.2f", num(d, "high")),
			fmt.Sprintf("%.2f", num(d, "low")),
			fmt.Sprintf("%.2f", num(d, "close")),
			str(d, "volume"),
		})
	}
	printTable(headers, rows)
}

// displayTimeSales renders time and sales data as a table.
func displayTimeSales(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	s := nested(root, "series")
	if s == nil {
		fmt.Println("No time and sales data found.")
		return
	}

	entries := toSlice(s["data"])
	headers := []string{"TIMESTAMP", "OPEN", "HIGH", "LOW", "CLOSE", "VOLUME"}
	rows := make([][]string, 0, len(entries))
	for _, e := range entries {
		rows = append(rows, []string{
			str(e, "timestamp"),
			fmt.Sprintf("%.2f", num(e, "open")),
			fmt.Sprintf("%.2f", num(e, "high")),
			fmt.Sprintf("%.2f", num(e, "low")),
			fmt.Sprintf("%.2f", num(e, "close")),
			str(e, "volume"),
		})
	}
	printTable(headers, rows)
}

// displayCalendar renders the market calendar as a table.
func displayCalendar(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	cal := nested(root, "calendar")
	if cal == nil {
		fmt.Println("No calendar data found.")
		return
	}

	days := nested(cal, "days")
	if days == nil {
		fmt.Println("No calendar days found.")
		return
	}

	dayList := toSlice(days["day"])
	headers := []string{"DATE", "STATUS", "DESCRIPTION", "OPEN", "CLOSE"}
	rows := make([][]string, 0, len(dayList))
	for _, d := range dayList {
		open := str(nested(d, "open"), "start")
		close := str(nested(d, "open"), "end")
		rows = append(rows, []string{
			shortDate(str(d, "date")),
			str(d, "status"),
			str(d, "description"),
			open,
			close,
		})
	}
	printTable(headers, rows)
}

// displayClock renders the current market clock status as key-value pairs.
func displayClock(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	c := nested(root, "clock")
	if c == nil {
		fmt.Println("No clock data found.")
		return
	}

	pairs := [][2]string{
		{"Date", str(c, "date")},
		{"State", str(c, "state")},
		{"Description", str(c, "description")},
		{"Next State", str(c, "next_state")},
		{"Next Change", str(c, "next_change")},
	}
	printKV(pairs)
}

// displayETB renders the easy-to-borrow securities list as a table.
func displayETB(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	sec := nested(root, "securities")
	if sec == nil {
		fmt.Println("No ETB data found.")
		return
	}

	securities := toSlice(sec["security"])
	headers := []string{"SYMBOL"}
	rows := make([][]string, 0, len(securities))
	for _, s := range securities {
		rows = append(rows, []string{str(s, "symbol")})
	}
	printTable(headers, rows)
}

// displaySecurities renders market lookup/search results as a table.
func displaySecurities(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	sec := nested(root, "securities")
	if sec == nil {
		fmt.Println("No securities found.")
		return
	}

	securities := toSlice(sec["security"])
	headers := []string{"SYMBOL", "EXCHANGE", "TYPE", "DESCRIPTION"}
	rows := make([][]string, 0, len(securities))
	for _, s := range securities {
		rows = append(rows, []string{
			str(s, "symbol"),
			str(s, "exchange"),
			str(s, "type"),
			str(s, "description"),
		})
	}
	printTable(headers, rows)
}

// ===========================================================================
// Trading Display Functions
// ===========================================================================

// displayOrderResult renders an order placement/modification/cancellation result.
func displayOrderResult(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	o := nested(root, "order")
	if o == nil {
		fmt.Println(string(data))
		return
	}

	pairs := [][2]string{
		{"Order ID", str(o, "id")},
		{"Status", str(o, "status")},
	}

	if partnerID := str(o, "partner_id"); partnerID != "" {
		pairs = append(pairs, [2]string{"Partner ID", partnerID})
	}

	printKV(pairs)
}

// ===========================================================================
// User Display Functions
// ===========================================================================

// displayProfile renders user profile information with associated accounts.
func displayProfile(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	p := nested(root, "profile")
	if p == nil {
		fmt.Println("No profile data found.")
		return
	}

	fmt.Printf("User: %s (%s)\n\n", str(p, "name"), str(p, "id"))

	accounts := toSlice(p["account"])
	if len(accounts) > 0 {
		headers := []string{"ACCOUNT", "CLASSIFICATION", "STATUS", "TYPE", "OPTION LEVEL", "DAY TRADER"}
		rows := make([][]string, 0, len(accounts))
		for _, a := range accounts {
			rows = append(rows, []string{
				str(a, "account_number"),
				str(a, "classification"),
				str(a, "status"),
				str(a, "type"),
				str(a, "option_level"),
				str(a, "day_trader"),
			})
		}
		printTable(headers, rows)
	}
}

// ===========================================================================
// Watchlist Display Functions
// ===========================================================================

// displayWatchlists renders all watchlists as a table.
func displayWatchlists(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	wl := nested(root, "watchlists")
	if wl == nil {
		fmt.Println("No watchlists found.")
		return
	}

	lists := toSlice(wl["watchlist"])
	headers := []string{"ID", "NAME"}
	rows := make([][]string, 0, len(lists))
	for _, l := range lists {
		rows = append(rows, []string{
			str(l, "id"),
			str(l, "name"),
		})
	}
	printTable(headers, rows)
}

// displayWatchlist renders a single watchlist with its symbols.
func displayWatchlist(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	wl := nested(root, "watchlist")
	if wl == nil {
		// Might be a direct response
		if str(root, "id") != "" {
			wl = root
		} else {
			fmt.Println("Watchlist updated.")
			return
		}
	}

	fmt.Printf("Watchlist: %s (%s)\n", str(wl, "name"), str(wl, "id"))

	items := nested(wl, "items")
	if items == nil {
		return
	}

	symbols := toSlice(items["item"])
	if len(symbols) > 0 {
		fmt.Println()
		headers := []string{"SYMBOL"}
		rows := make([][]string, 0, len(symbols))
		for _, s := range symbols {
			rows = append(rows, []string{str(s, "symbol")})
		}
		printTable(headers, rows)
	}
}

// ===========================================================================
// Streaming Display Functions
// ===========================================================================

// displaySession renders a streaming session result.
func displaySession(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	s := nested(root, "stream")
	if s == nil {
		fmt.Println(string(data))
		return
	}

	pairs := [][2]string{
		{"Session ID", str(s, "sessionid")},
		{"URL", str(s, "url")},
	}
	printKV(pairs)
}

// ===========================================================================
// Generic Display Function
// ===========================================================================

// displayGeneric prints a simple status message from a generic API response.
func displayGeneric(data []byte) {
	root := parseJSON(data)
	if root == nil {
		fmt.Println(string(data))
		return
	}

	// Try to find a status or display something useful
	if status := str(root, "status"); status != "" {
		fmt.Printf("Status: %s\n", status)
		return
	}

	// Fall back to key-value display of top-level fields
	pairs := make([][2]string, 0)
	for k, v := range root {
		pairs = append(pairs, [2]string{k, fmt.Sprintf("%v", v)})
	}
	if len(pairs) > 0 {
		printKV(pairs)
	} else {
		fmt.Println("OK")
	}
}
