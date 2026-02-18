// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

// GetProfile retrieves the profile information for the currently authenticated user.
func (c *Client) GetProfile() ([]byte, error) {
	return c.doGet("/v1/user/profile", nil)
}
