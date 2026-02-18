// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetBalances verifies the GetBalances method sends correct request and returns response body.
func TestGetBalances(t *testing.T) {
	body := `{"balances":{"total_equity":17798.36,"account_number":"VA000001"}}`
	server := testServer(t, "GET", "/v1/accounts/VA000001/balances", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetBalances("VA000001")
	if err != nil {
		t.Fatalf("GetBalances() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetBalances() = %s, want %s", result, body)
	}
}

// TestGetGainLoss verifies the GetGainLoss method sends correct request with query parameters.
func TestGetGainLoss(t *testing.T) {
	body := `{"gainloss":{"closed_position":[]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/accounts/VA000001/gainloss" {
			t.Errorf("path = %s, want /v1/accounts/VA000001/gainloss", r.URL.Path)
		}
		if r.URL.Query().Get("page") != "1" {
			t.Errorf("page = %s, want 1", r.URL.Query().Get("page"))
		}
		if r.URL.Query().Get("limit") != "50" {
			t.Errorf("limit = %s, want 50", r.URL.Query().Get("limit"))
		}
		if r.URL.Query().Get("sortBy") != "closedate" {
			t.Errorf("sortBy = %s, want closedate", r.URL.Query().Get("sortBy"))
		}
		if r.URL.Query().Get("sort") != "desc" {
			t.Errorf("sort = %s, want desc", r.URL.Query().Get("sort"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetGainLoss("VA000001", "1", "50", "closedate", "desc")
	if err != nil {
		t.Fatalf("GetGainLoss() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetGainLoss() = %s, want %s", result, body)
	}
}

// TestGetHistoricalBalances verifies the GetHistoricalBalances method with a period parameter.
func TestGetHistoricalBalances(t *testing.T) {
	body := `{"balances":[{"date":"2024-01-01","value":10000}]}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/accounts/VA000001/historical-balances" {
			t.Errorf("path = %s, want /v1/accounts/VA000001/historical-balances", r.URL.Path)
		}
		if r.URL.Query().Get("period") != "MONTH" {
			t.Errorf("period = %s, want MONTH", r.URL.Query().Get("period"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetHistoricalBalances("VA000001", "MONTH")
	if err != nil {
		t.Fatalf("GetHistoricalBalances() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetHistoricalBalances() = %s, want %s", result, body)
	}
}

// TestGetHistory verifies the GetHistory method with pagination and type filter parameters.
func TestGetHistory(t *testing.T) {
	body := `{"history":{"event":[]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/accounts/VA000001/history" {
			t.Errorf("path = %s, want /v1/accounts/VA000001/history", r.URL.Path)
		}
		if r.URL.Query().Get("page") != "2" {
			t.Errorf("page = %s, want 2", r.URL.Query().Get("page"))
		}
		if r.URL.Query().Get("limit") != "25" {
			t.Errorf("limit = %s, want 25", r.URL.Query().Get("limit"))
		}
		if r.URL.Query().Get("type") != "trade" {
			t.Errorf("type = %s, want trade", r.URL.Query().Get("type"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetHistory("VA000001", "2", "25", "trade", "", "")
	if err != nil {
		t.Fatalf("GetHistory() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetHistory() = %s, want %s", result, body)
	}
}

// TestGetOrder verifies the GetOrder method with the correct path including order ID.
func TestGetOrder(t *testing.T) {
	body := `{"order":{"id":123456,"status":"filled"}}`
	server := testServer(t, "GET", "/v1/accounts/VA000001/orders/123456", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetOrder("VA000001", "123456", "")
	if err != nil {
		t.Fatalf("GetOrder() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetOrder() = %s, want %s", result, body)
	}
}

// TestGetOrders verifies the GetOrders method returns all orders for an account.
func TestGetOrders(t *testing.T) {
	body := `{"orders":{"order":[{"id":1},{"id":2}]}}`
	server := testServer(t, "GET", "/v1/accounts/VA000001/orders", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetOrders("VA000001", "", "", "")
	if err != nil {
		t.Fatalf("GetOrders() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetOrders() = %s, want %s", result, body)
	}
}

// TestGetPositions verifies the GetPositions method returns current positions.
func TestGetPositions(t *testing.T) {
	body := `{"positions":{"position":[{"symbol":"AAPL","quantity":37}]}}`
	server := testServer(t, "GET", "/v1/accounts/VA000001/positions", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetPositions("VA000001")
	if err != nil {
		t.Fatalf("GetPositions() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetPositions() = %s, want %s", result, body)
	}
}

// TestGetPositionGroups verifies the GetPositionGroups method returns all groups.
func TestGetPositionGroups(t *testing.T) {
	body := `{"position_groups":[{"id":1,"label":"Tech"}]}`
	server := testServer(t, "GET", "/v1/accounts/VA000001/position-groups", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetPositionGroups("VA000001")
	if err != nil {
		t.Fatalf("GetPositionGroups() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetPositionGroups() = %s, want %s", result, body)
	}
}

// TestCreatePositionGroup verifies the CreatePositionGroup method sends a POST with form data.
func TestCreatePositionGroup(t *testing.T) {
	body := `{"position_group":{"id":1,"label":"Big Tech"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/accounts/VA000001/position-groups" {
			t.Errorf("path = %s, want /v1/accounts/VA000001/position-groups", r.URL.Path)
		}
		r.ParseForm()
		if r.PostForm.Get("label") != "Big Tech" {
			t.Errorf("label = %s, want Big Tech", r.PostForm.Get("label"))
		}
		if r.PostForm.Get("symbols") != "AAPL,MSFT" {
			t.Errorf("symbols = %s, want AAPL,MSFT", r.PostForm.Get("symbols"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.CreatePositionGroup("VA000001", "Big Tech", "AAPL,MSFT")
	if err != nil {
		t.Fatalf("CreatePositionGroup() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("CreatePositionGroup() = %s, want %s", result, body)
	}
}

// TestUpdatePositionGroup verifies the UpdatePositionGroup method sends a PUT request.
func TestUpdatePositionGroup(t *testing.T) {
	body := `{"position_group":{"id":1,"label":"Updated Tech"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		if r.URL.Path != "/v1/accounts/VA000001/position-groups/1" {
			t.Errorf("path = %s, want /v1/accounts/VA000001/position-groups/1", r.URL.Path)
		}
		r.ParseForm()
		if r.PostForm.Get("label") != "Updated Tech" {
			t.Errorf("label = %s, want Updated Tech", r.PostForm.Get("label"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.UpdatePositionGroup("VA000001", "1", "Updated Tech", "AAPL,GOOGL")
	if err != nil {
		t.Fatalf("UpdatePositionGroup() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("UpdatePositionGroup() = %s, want %s", result, body)
	}
}

// TestDeletePositionGroup verifies the DeletePositionGroup method sends a DELETE request.
func TestDeletePositionGroup(t *testing.T) {
	body := `{"status":"ok"}`
	server := testServer(t, "DELETE", "/v1/accounts/VA000001/position-groups/1", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.DeletePositionGroup("VA000001", "1")
	if err != nil {
		t.Fatalf("DeletePositionGroup() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("DeletePositionGroup() = %s, want %s", result, body)
	}
}
