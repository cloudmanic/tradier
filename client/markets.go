// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

// GetQuotes retrieves real-time quotes for one or more symbols (comma-separated).
func (c *Client) GetQuotes(symbols string, greeks string) ([]byte, error) {
	params := map[string]string{
		"symbols": symbols,
		"greeks":  greeks,
	}
	return c.doGet("/v1/markets/quotes", params)
}

// PostQuotes retrieves quotes for a larger list of symbols via POST request.
func (c *Client) PostQuotes(symbols string, greeks string) ([]byte, error) {
	params := map[string]string{
		"symbols": symbols,
		"greeks":  greeks,
	}
	return c.doPost("/v1/markets/quotes", params)
}

// GetOptionsChains retrieves option chains for a specific symbol and expiration date.
func (c *Client) GetOptionsChains(symbol, expiration, greeks string) ([]byte, error) {
	params := map[string]string{
		"symbol":     symbol,
		"expiration": expiration,
		"greeks":     greeks,
	}
	return c.doGet("/v1/markets/options/chains", params)
}

// GetOptionsExpirations retrieves available expiration dates for a specific underlying symbol.
func (c *Client) GetOptionsExpirations(symbol, includeAllRoots, strikes, contractSize, expirationType string) ([]byte, error) {
	params := map[string]string{
		"symbol":          symbol,
		"includeAllRoots": includeAllRoots,
		"strikes":         strikes,
		"contractSize":    contractSize,
		"expirationType":  expirationType,
	}
	return c.doGet("/v1/markets/options/expirations", params)
}

// GetOptionsStrikes retrieves available strike prices for a specific symbol and expiration date.
func (c *Client) GetOptionsStrikes(symbol, expiration string) ([]byte, error) {
	params := map[string]string{
		"symbol":     symbol,
		"expiration": expiration,
	}
	return c.doGet("/v1/markets/options/strikes", params)
}

// GetOptionsLookup retrieves all options symbols for a given underlying with optional filters.
func (c *Client) GetOptionsLookup(underlying, strike, expiration, optionType string) ([]byte, error) {
	params := map[string]string{
		"underlying": underlying,
		"strike":     strike,
		"expiration": expiration,
		"type":       optionType,
	}
	return c.doGet("/v1/markets/options/lookup", params)
}

// GetHistoricalPricing retrieves historical OHLCV pricing data for a security.
func (c *Client) GetHistoricalPricing(symbol, interval, start, end string) ([]byte, error) {
	params := map[string]string{
		"symbol":   symbol,
		"interval": interval,
		"start":    start,
		"end":      end,
	}
	return c.doGet("/v1/markets/history", params)
}

// GetTimeSales retrieves time and sales data for charting at predefined intervals.
func (c *Client) GetTimeSales(symbol, interval, start, end, sessionFilter string) ([]byte, error) {
	params := map[string]string{
		"symbol":         symbol,
		"interval":       interval,
		"start":          start,
		"end":            end,
		"session_filter": sessionFilter,
	}
	return c.doGet("/v1/markets/timesales", params)
}

// GetCalendar retrieves the market calendar for the current or a specific month/year.
func (c *Client) GetCalendar(month, year string) ([]byte, error) {
	params := map[string]string{
		"month": month,
		"year":  year,
	}
	return c.doGet("/v1/markets/calendar", params)
}

// GetClock retrieves the current intraday market status (pre, open, post, closed).
func (c *Client) GetClock() ([]byte, error) {
	return c.doGet("/v1/markets/clock", nil)
}

// GetETB retrieves the list of Easy-To-Borrow securities available for short selling.
func (c *Client) GetETB() ([]byte, error) {
	return c.doGet("/v1/markets/etb", nil)
}

// GetLookup searches for a symbol using the ticker symbol or partial symbol.
func (c *Client) GetLookup(query, exchanges, types string) ([]byte, error) {
	params := map[string]string{
		"q":         query,
		"exchanges": exchanges,
		"types":     types,
	}
	return c.doGet("/v1/markets/lookup", params)
}

// GetSearch searches for securities by partial match on symbol or company name.
func (c *Client) GetSearch(query, indexes string) ([]byte, error) {
	params := map[string]string{
		"q":       query,
		"indexes": indexes,
	}
	return c.doGet("/v1/markets/search", params)
}
