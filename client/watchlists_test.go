// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetWatchlists verifies the GetWatchlists method returns all watchlists.
func TestGetWatchlists(t *testing.T) {
	body := `{"watchlists":{"watchlist":[{"name":"My List","id":"wl-123"}]}}`
	server := testServer(t, "GET", "/v1/watchlists", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetWatchlists()
	if err != nil {
		t.Fatalf("GetWatchlists() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetWatchlists() = %s, want %s", result, body)
	}
}

// TestGetWatchlist verifies the GetWatchlist method retrieves a specific watchlist by ID.
func TestGetWatchlist(t *testing.T) {
	body := `{"watchlist":{"name":"My List","id":"wl-123","items":{"item":[{"symbol":"AAPL"}]}}}`
	server := testServer(t, "GET", "/v1/watchlists/wl-123", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetWatchlist("wl-123")
	if err != nil {
		t.Fatalf("GetWatchlist() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetWatchlist() = %s, want %s", result, body)
	}
}

// TestCreateWatchlist verifies the CreateWatchlist method sends a POST with form data.
func TestCreateWatchlist(t *testing.T) {
	body := `{"watchlist":{"name":"Tech Stocks","id":"wl-456"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/watchlists" {
			t.Errorf("path = %s, want /v1/watchlists", r.URL.Path)
		}
		r.ParseForm()
		if r.PostForm.Get("name") != "Tech Stocks" {
			t.Errorf("name = %s, want Tech Stocks", r.PostForm.Get("name"))
		}
		if r.PostForm.Get("symbols") != "AAPL,MSFT,GOOGL" {
			t.Errorf("symbols = %s, want AAPL,MSFT,GOOGL", r.PostForm.Get("symbols"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.CreateWatchlist("Tech Stocks", "AAPL,MSFT,GOOGL")
	if err != nil {
		t.Fatalf("CreateWatchlist() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("CreateWatchlist() = %s, want %s", result, body)
	}
}

// TestUpdateWatchlist verifies the UpdateWatchlist method sends a PUT with form data.
func TestUpdateWatchlist(t *testing.T) {
	body := `{"watchlist":{"name":"Updated List","id":"wl-123"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		if r.URL.Path != "/v1/watchlists/wl-123" {
			t.Errorf("path = %s, want /v1/watchlists/wl-123", r.URL.Path)
		}
		r.ParseForm()
		if r.PostForm.Get("name") != "Updated List" {
			t.Errorf("name = %s, want Updated List", r.PostForm.Get("name"))
		}
		if r.PostForm.Get("symbols") != "AAPL,TSLA" {
			t.Errorf("symbols = %s, want AAPL,TSLA", r.PostForm.Get("symbols"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.UpdateWatchlist("wl-123", "Updated List", "AAPL,TSLA")
	if err != nil {
		t.Fatalf("UpdateWatchlist() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("UpdateWatchlist() = %s, want %s", result, body)
	}
}

// TestDeleteWatchlist verifies the DeleteWatchlist method sends a DELETE request.
func TestDeleteWatchlist(t *testing.T) {
	body := `{"watchlists":{"watchlist":[]}}`
	server := testServer(t, "DELETE", "/v1/watchlists/wl-123", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.DeleteWatchlist("wl-123")
	if err != nil {
		t.Fatalf("DeleteWatchlist() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("DeleteWatchlist() = %s, want %s", result, body)
	}
}

// TestAddSymbolsToWatchlist verifies the AddSymbolsToWatchlist method sends a POST with symbols.
func TestAddSymbolsToWatchlist(t *testing.T) {
	body := `{"watchlist":{"name":"My List","id":"wl-123"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/watchlists/wl-123/symbols" {
			t.Errorf("path = %s, want /v1/watchlists/wl-123/symbols", r.URL.Path)
		}
		r.ParseForm()
		if r.PostForm.Get("symbols") != "SPY,QQQ" {
			t.Errorf("symbols = %s, want SPY,QQQ", r.PostForm.Get("symbols"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.AddSymbolsToWatchlist("wl-123", "SPY,QQQ")
	if err != nil {
		t.Fatalf("AddSymbolsToWatchlist() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("AddSymbolsToWatchlist() = %s, want %s", result, body)
	}
}

// TestRemoveSymbolFromWatchlist verifies the RemoveSymbolFromWatchlist method sends a DELETE.
func TestRemoveSymbolFromWatchlist(t *testing.T) {
	body := `{"watchlist":{"name":"My List","id":"wl-123"}}`
	server := testServer(t, "DELETE", "/v1/watchlists/wl-123/symbols/SPY", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.RemoveSymbolFromWatchlist("wl-123", "SPY")
	if err != nil {
		t.Fatalf("RemoveSymbolFromWatchlist() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("RemoveSymbolFromWatchlist() = %s, want %s", result, body)
	}
}
