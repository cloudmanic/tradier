// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import "testing"

// TestGetProfile verifies the GetProfile method returns user profile data.
func TestGetProfile(t *testing.T) {
	body := `{"profile":{"id":"id-123","name":"John Doe","account":[{"account_number":"VA000001","status":"active"}]}}`
	server := testServer(t, "GET", "/v1/user/profile", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetProfile()
	if err != nil {
		t.Fatalf("GetProfile() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetProfile() = %s, want %s", result, body)
	}
}
