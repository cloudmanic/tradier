// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import "fmt"

// PlaceOrder places a trading order for a given account. The params map should contain
// all required form parameters for the order class (equity, option, multileg, combo, oto, oco, otoco).
// Common params: class, symbol, side, quantity, type, duration, price, stop, tag, preview.
// For multileg/combo orders use indexed params like option_symbol[0], side[0], quantity[0], etc.
func (c *Client) PlaceOrder(accountID string, params map[string]string) ([]byte, error) {
	path := fmt.Sprintf("/v1/accounts/%s/orders", accountID)
	return c.doPost(path, params)
}

// ChangeOrder modifies an existing order. Supports changing type, duration, price, and stop.
func (c *Client) ChangeOrder(accountID, orderID string, params map[string]string) ([]byte, error) {
	path := fmt.Sprintf("/v1/accounts/%s/orders/%s", accountID, orderID)
	return c.doPut(path, params)
}

// CancelOrder cancels an existing order by its ID.
func (c *Client) CancelOrder(accountID, orderID string) ([]byte, error) {
	path := fmt.Sprintf("/v1/accounts/%s/orders/%s", accountID, orderID)
	return c.doDelete(path)
}
