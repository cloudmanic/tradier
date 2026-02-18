// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

// CreateMarketSession creates a streaming session for real-time market data via WebSocket.
func (c *Client) CreateMarketSession() ([]byte, error) {
	return c.doPost("/v1/markets/events/session", nil)
}

// CreateAccountSession creates a streaming session for real-time account events via WebSocket.
func (c *Client) CreateAccountSession() ([]byte, error) {
	return c.doPost("/v1/accounts/events/session", nil)
}
