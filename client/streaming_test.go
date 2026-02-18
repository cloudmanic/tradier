// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateMarketSession verifies the CreateMarketSession method sends a POST and returns session data.
func TestCreateMarketSession(t *testing.T) {
	body := `{"stream":{"url":"wss://ws.tradier.com/v1/markets/events","sessionid":"sess-abc-123"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/markets/events/session" {
			t.Errorf("path = %s, want /v1/markets/events/session", r.URL.Path)
		}
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-key" {
			t.Errorf("Authorization = %q, want %q", auth, "Bearer test-key")
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.CreateMarketSession()
	if err != nil {
		t.Fatalf("CreateMarketSession() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("CreateMarketSession() = %s, want %s", result, body)
	}
}

// TestCreateAccountSession verifies the CreateAccountSession method sends a POST and returns session data.
func TestCreateAccountSession(t *testing.T) {
	body := `{"stream":{"url":"wss://ws.tradier.com/v1/accounts/events","sessionid":"sess-xyz-789"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/accounts/events/session" {
			t.Errorf("path = %s, want /v1/accounts/events/session", r.URL.Path)
		}
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-key" {
			t.Errorf("Authorization = %q, want %q", auth, "Bearer test-key")
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.CreateAccountSession()
	if err != nil {
		t.Fatalf("CreateAccountSession() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("CreateAccountSession() = %s, want %s", result, body)
	}
}
