// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import "fmt"

// GetBalances retrieves the current balance and margin information for a specific account.
func (c *Client) GetBalances(accountID string) ([]byte, error) {
	path := fmt.Sprintf("/v1/accounts/%s/balances", accountID)
	return c.doGet(path, nil)
}

// GetGainLoss retrieves cost basis and gain/loss information for all closed positions in an account.
func (c *Client) GetGainLoss(accountID string, page, limit, sortBy, sort string) ([]byte, error) {
	params := map[string]string{
		"page":   page,
		"limit":  limit,
		"sortBy": sortBy,
		"sort":   sort,
	}
	path := fmt.Sprintf("/v1/accounts/%s/gainloss", accountID)
	return c.doGet(path, params)
}

// GetHistoricalBalances retrieves historical account balances to track value over time.
func (c *Client) GetHistoricalBalances(accountID, period string) ([]byte, error) {
	params := map[string]string{
		"period": period,
	}
	path := fmt.Sprintf("/v1/accounts/%s/historical-balances", accountID)
	return c.doGet(path, params)
}

// GetHistory retrieves historical activity events for an account with optional filtering.
func (c *Client) GetHistory(accountID, page, limit, activityType, start, end string) ([]byte, error) {
	params := map[string]string{
		"page":  page,
		"limit": limit,
		"type":  activityType,
		"start": start,
		"end":   end,
	}
	path := fmt.Sprintf("/v1/accounts/%s/history", accountID)
	return c.doGet(path, params)
}

// GetOrder retrieves a specific order by its ID for a given account.
func (c *Client) GetOrder(accountID string, orderID string, includeTags string) ([]byte, error) {
	params := map[string]string{
		"includeTags": includeTags,
	}
	path := fmt.Sprintf("/v1/accounts/%s/orders/%s", accountID, orderID)
	return c.doGet(path, params)
}

// GetOrders retrieves all orders for a given account with optional pagination.
func (c *Client) GetOrders(accountID, page, limit, includeTags string) ([]byte, error) {
	params := map[string]string{
		"page":        page,
		"limit":       limit,
		"includeTags": includeTags,
	}
	path := fmt.Sprintf("/v1/accounts/%s/orders", accountID)
	return c.doGet(path, params)
}

// GetPositions retrieves the current positions held in an account.
func (c *Client) GetPositions(accountID string) ([]byte, error) {
	path := fmt.Sprintf("/v1/accounts/%s/positions", accountID)
	return c.doGet(path, nil)
}

// GetPositionGroups retrieves all position groups for a specific account.
func (c *Client) GetPositionGroups(accountID string) ([]byte, error) {
	path := fmt.Sprintf("/v1/accounts/%s/position-groups", accountID)
	return c.doGet(path, nil)
}

// CreatePositionGroup creates a new position group for a specific account with the given label and symbols.
func (c *Client) CreatePositionGroup(accountID, label, symbols string) ([]byte, error) {
	params := map[string]string{
		"label":   label,
		"symbols": symbols,
	}
	path := fmt.Sprintf("/v1/accounts/%s/position-groups", accountID)
	return c.doPost(path, params)
}

// UpdatePositionGroup updates an existing position group with a new label and symbols.
func (c *Client) UpdatePositionGroup(accountID, groupID, label, symbols string) ([]byte, error) {
	params := map[string]string{
		"label":   label,
		"symbols": symbols,
	}
	path := fmt.Sprintf("/v1/accounts/%s/position-groups/%s", accountID, groupID)
	return c.doPut(path, params)
}

// DeletePositionGroup deletes a position group from a specific account.
func (c *Client) DeletePositionGroup(accountID, groupID string) ([]byte, error) {
	path := fmt.Sprintf("/v1/accounts/%s/position-groups/%s", accountID, groupID)
	return c.doDelete(path)
}
