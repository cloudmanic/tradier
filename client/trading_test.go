// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPlaceEquityOrder verifies placing a simple equity market order via POST.
func TestPlaceEquityOrder(t *testing.T) {
	body := `{"order":{"id":257459,"status":"ok","partner_id":"1"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/accounts/VA000001/orders" {
			t.Errorf("path = %s, want /v1/accounts/VA000001/orders", r.URL.Path)
		}
		r.ParseForm()
		if r.PostForm.Get("class") != "equity" {
			t.Errorf("class = %s, want equity", r.PostForm.Get("class"))
		}
		if r.PostForm.Get("symbol") != "AAPL" {
			t.Errorf("symbol = %s, want AAPL", r.PostForm.Get("symbol"))
		}
		if r.PostForm.Get("side") != "buy" {
			t.Errorf("side = %s, want buy", r.PostForm.Get("side"))
		}
		if r.PostForm.Get("quantity") != "10" {
			t.Errorf("quantity = %s, want 10", r.PostForm.Get("quantity"))
		}
		if r.PostForm.Get("type") != "market" {
			t.Errorf("type = %s, want market", r.PostForm.Get("type"))
		}
		if r.PostForm.Get("duration") != "day" {
			t.Errorf("duration = %s, want day", r.PostForm.Get("duration"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	params := map[string]string{
		"class":    "equity",
		"symbol":   "AAPL",
		"side":     "buy",
		"quantity": "10",
		"type":     "market",
		"duration": "day",
	}
	result, err := c.PlaceOrder("VA000001", params)
	if err != nil {
		t.Fatalf("PlaceOrder() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("PlaceOrder() = %s, want %s", result, body)
	}
}

// TestPlaceOptionOrder verifies placing an option order via POST.
func TestPlaceOptionOrder(t *testing.T) {
	body := `{"order":{"id":257460,"status":"ok"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		r.ParseForm()
		if r.PostForm.Get("class") != "option" {
			t.Errorf("class = %s, want option", r.PostForm.Get("class"))
		}
		if r.PostForm.Get("symbol") != "AAPL" {
			t.Errorf("symbol = %s, want AAPL", r.PostForm.Get("symbol"))
		}
		if r.PostForm.Get("option_symbol") != "AAPL220617C00270000" {
			t.Errorf("option_symbol = %s, want AAPL220617C00270000", r.PostForm.Get("option_symbol"))
		}
		if r.PostForm.Get("side") != "buy_to_open" {
			t.Errorf("side = %s, want buy_to_open", r.PostForm.Get("side"))
		}
		if r.PostForm.Get("quantity") != "5" {
			t.Errorf("quantity = %s, want 5", r.PostForm.Get("quantity"))
		}
		if r.PostForm.Get("type") != "limit" {
			t.Errorf("type = %s, want limit", r.PostForm.Get("type"))
		}
		if r.PostForm.Get("price") != "3.50" {
			t.Errorf("price = %s, want 3.50", r.PostForm.Get("price"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	params := map[string]string{
		"class":         "option",
		"symbol":        "AAPL",
		"option_symbol": "AAPL220617C00270000",
		"side":          "buy_to_open",
		"quantity":      "5",
		"type":          "limit",
		"duration":      "day",
		"price":         "3.50",
	}
	result, err := c.PlaceOrder("VA000001", params)
	if err != nil {
		t.Fatalf("PlaceOrder() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("PlaceOrder() = %s, want %s", result, body)
	}
}

// TestPlaceMultilegOrder verifies placing a multileg options order via POST.
func TestPlaceMultilegOrder(t *testing.T) {
	body := `{"order":{"id":257461,"status":"ok"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		r.ParseForm()
		if r.PostForm.Get("class") != "multileg" {
			t.Errorf("class = %s, want multileg", r.PostForm.Get("class"))
		}
		if r.PostForm.Get("symbol") != "AAPL" {
			t.Errorf("symbol = %s, want AAPL", r.PostForm.Get("symbol"))
		}
		if r.PostForm.Get("type") != "debit" {
			t.Errorf("type = %s, want debit", r.PostForm.Get("type"))
		}
		if r.PostForm.Get("option_symbol[0]") != "AAPL220617C00270000" {
			t.Errorf("option_symbol[0] = %s, want AAPL220617C00270000", r.PostForm.Get("option_symbol[0]"))
		}
		if r.PostForm.Get("side[0]") != "buy_to_open" {
			t.Errorf("side[0] = %s, want buy_to_open", r.PostForm.Get("side[0]"))
		}
		if r.PostForm.Get("quantity[0]") != "1" {
			t.Errorf("quantity[0] = %s, want 1", r.PostForm.Get("quantity[0]"))
		}
		if r.PostForm.Get("option_symbol[1]") != "AAPL220617C00280000" {
			t.Errorf("option_symbol[1] = %s, want AAPL220617C00280000", r.PostForm.Get("option_symbol[1]"))
		}
		if r.PostForm.Get("side[1]") != "sell_to_open" {
			t.Errorf("side[1] = %s, want sell_to_open", r.PostForm.Get("side[1]"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	params := map[string]string{
		"class":            "multileg",
		"symbol":           "AAPL",
		"type":             "debit",
		"duration":         "day",
		"price":            "1.50",
		"option_symbol[0]": "AAPL220617C00270000",
		"side[0]":          "buy_to_open",
		"quantity[0]":      "1",
		"option_symbol[1]": "AAPL220617C00280000",
		"side[1]":          "sell_to_open",
		"quantity[1]":      "1",
	}
	result, err := c.PlaceOrder("VA000001", params)
	if err != nil {
		t.Fatalf("PlaceOrder() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("PlaceOrder() = %s, want %s", result, body)
	}
}

// TestChangeOrder verifies that ChangeOrder sends a PUT request with form parameters.
func TestChangeOrder(t *testing.T) {
	body := `{"order":{"id":123456,"status":"ok"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		if r.URL.Path != "/v1/accounts/VA000001/orders/123456" {
			t.Errorf("path = %s, want /v1/accounts/VA000001/orders/123456", r.URL.Path)
		}
		r.ParseForm()
		if r.PostForm.Get("type") != "limit" {
			t.Errorf("type = %s, want limit", r.PostForm.Get("type"))
		}
		if r.PostForm.Get("price") != "155.00" {
			t.Errorf("price = %s, want 155.00", r.PostForm.Get("price"))
		}
		if r.PostForm.Get("duration") != "gtc" {
			t.Errorf("duration = %s, want gtc", r.PostForm.Get("duration"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	params := map[string]string{
		"type":     "limit",
		"price":    "155.00",
		"duration": "gtc",
	}
	result, err := c.ChangeOrder("VA000001", "123456", params)
	if err != nil {
		t.Fatalf("ChangeOrder() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("ChangeOrder() = %s, want %s", result, body)
	}
}

// TestCancelOrder verifies that CancelOrder sends a DELETE request.
func TestCancelOrder(t *testing.T) {
	body := `{"order":{"id":123456,"status":"ok"}}`
	server := testServer(t, "DELETE", "/v1/accounts/VA000001/orders/123456", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.CancelOrder("VA000001", "123456")
	if err != nil {
		t.Fatalf("CancelOrder() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("CancelOrder() = %s, want %s", result, body)
	}
}
