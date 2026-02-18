// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// testServer creates a mock HTTP server that verifies the request method, path, and auth header,
// then responds with the given status code and body.
func testServer(t *testing.T, wantMethod, wantPath string, statusCode int, responseBody string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != wantMethod {
			t.Errorf("method = %s, want %s", r.Method, wantMethod)
		}
		if r.URL.Path != wantPath {
			t.Errorf("path = %s, want %s", r.URL.Path, wantPath)
		}
		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-key" {
			t.Errorf("Authorization = %q, want %q", auth, "Bearer test-key")
		}
		accept := r.Header.Get("Accept")
		if accept != "application/json" {
			t.Errorf("Accept = %q, want %q", accept, "application/json")
		}
		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	}))
}

// testClient creates a Client configured to point at the given test server.
func testClient(server *httptest.Server) *Client {
	return NewClient(server.URL, "test-key")
}

// TestNewClient verifies that NewClient sets fields correctly.
func TestNewClient(t *testing.T) {
	c := NewClient("https://api.tradier.com", "my-key")
	if c.BaseURL != "https://api.tradier.com" {
		t.Errorf("BaseURL = %q, want %q", c.BaseURL, "https://api.tradier.com")
	}
	if c.APIKey != "my-key" {
		t.Errorf("APIKey = %q, want %q", c.APIKey, "my-key")
	}
	if c.HTTPClient == nil {
		t.Error("HTTPClient should not be nil")
	}
}

// TestDoGetErrorStatus verifies that non-2xx status codes return an error.
func TestDoGetErrorStatus(t *testing.T) {
	server := testServer(t, "GET", "/v1/test", 401, `{"error":"unauthorized"}`)
	defer server.Close()
	c := testClient(server)

	_, err := c.doGet("/v1/test", nil)
	if err == nil {
		t.Fatal("expected error for 401 status")
	}
}

// TestPrettyJSON verifies that PrettyJSON formats JSON correctly.
func TestPrettyJSON(t *testing.T) {
	input := []byte(`{"key":"value","num":42}`)
	result, err := PrettyJSON(input)
	if err != nil {
		t.Fatalf("PrettyJSON error: %v", err)
	}
	expected := "{\n  \"key\": \"value\",\n  \"num\": 42\n}"
	if result != expected {
		t.Errorf("PrettyJSON result = %q, want %q", result, expected)
	}
}
