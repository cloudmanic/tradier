// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import "fmt"

// GetWatchlists retrieves all watchlists for the authenticated user.
func (c *Client) GetWatchlists() ([]byte, error) {
	return c.doGet("/v1/watchlists", nil)
}

// GetWatchlist retrieves a specific watchlist by its ID.
func (c *Client) GetWatchlist(watchlistID string) ([]byte, error) {
	path := fmt.Sprintf("/v1/watchlists/%s", watchlistID)
	return c.doGet(path, nil)
}

// CreateWatchlist creates a new watchlist with the specified name and comma-separated symbols.
func (c *Client) CreateWatchlist(name, symbols string) ([]byte, error) {
	params := map[string]string{
		"name":    name,
		"symbols": symbols,
	}
	return c.doPost("/v1/watchlists", params)
}

// UpdateWatchlist updates an existing watchlist with a new name and optional symbols.
func (c *Client) UpdateWatchlist(watchlistID, name, symbols string) ([]byte, error) {
	params := map[string]string{
		"name":    name,
		"symbols": symbols,
	}
	path := fmt.Sprintf("/v1/watchlists/%s", watchlistID)
	return c.doPut(path, params)
}

// DeleteWatchlist deletes a specific watchlist by its ID.
func (c *Client) DeleteWatchlist(watchlistID string) ([]byte, error) {
	path := fmt.Sprintf("/v1/watchlists/%s", watchlistID)
	return c.doDelete(path)
}

// AddSymbolsToWatchlist adds comma-separated symbols to an existing watchlist.
func (c *Client) AddSymbolsToWatchlist(watchlistID, symbols string) ([]byte, error) {
	params := map[string]string{
		"symbols": symbols,
	}
	path := fmt.Sprintf("/v1/watchlists/%s/symbols", watchlistID)
	return c.doPost(path, params)
}

// RemoveSymbolFromWatchlist removes a single symbol from a specific watchlist.
func (c *Client) RemoveSymbolFromWatchlist(watchlistID, symbol string) ([]byte, error) {
	path := fmt.Sprintf("/v1/watchlists/%s/symbols/%s", watchlistID, symbol)
	return c.doDelete(path)
}
