// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetQuotes verifies the GetQuotes method sends the correct symbols parameter.
func TestGetQuotes(t *testing.T) {
	body := `{"quotes":{"quote":{"symbol":"AAPL","last":150.00}}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if r.URL.Path != "/v1/markets/quotes" {
			t.Errorf("path = %s, want /v1/markets/quotes", r.URL.Path)
		}
		if r.URL.Query().Get("symbols") != "AAPL,SPY" {
			t.Errorf("symbols = %s, want AAPL,SPY", r.URL.Query().Get("symbols"))
		}
		if r.URL.Query().Get("greeks") != "true" {
			t.Errorf("greeks = %s, want true", r.URL.Query().Get("greeks"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetQuotes("AAPL,SPY", "true")
	if err != nil {
		t.Fatalf("GetQuotes() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetQuotes() = %s, want %s", result, body)
	}
}

// TestPostQuotes verifies the PostQuotes method sends a POST request with form data.
func TestPostQuotes(t *testing.T) {
	body := `{"quotes":{"quote":[{"symbol":"AAPL"},{"symbol":"MSFT"}]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/v1/markets/quotes" {
			t.Errorf("path = %s, want /v1/markets/quotes", r.URL.Path)
		}
		r.ParseForm()
		if r.PostForm.Get("symbols") != "AAPL,MSFT,GOOGL" {
			t.Errorf("symbols = %s, want AAPL,MSFT,GOOGL", r.PostForm.Get("symbols"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.PostQuotes("AAPL,MSFT,GOOGL", "")
	if err != nil {
		t.Fatalf("PostQuotes() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("PostQuotes() = %s, want %s", result, body)
	}
}

// TestGetOptionsChains verifies the GetOptionsChains method with symbol and expiration.
func TestGetOptionsChains(t *testing.T) {
	body := `{"options":{"option":[{"symbol":"AAPL220617C00270000"}]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/markets/options/chains" {
			t.Errorf("path = %s, want /v1/markets/options/chains", r.URL.Path)
		}
		if r.URL.Query().Get("symbol") != "AAPL" {
			t.Errorf("symbol = %s, want AAPL", r.URL.Query().Get("symbol"))
		}
		if r.URL.Query().Get("expiration") != "2024-06-21" {
			t.Errorf("expiration = %s, want 2024-06-21", r.URL.Query().Get("expiration"))
		}
		if r.URL.Query().Get("greeks") != "true" {
			t.Errorf("greeks = %s, want true", r.URL.Query().Get("greeks"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetOptionsChains("AAPL", "2024-06-21", "true")
	if err != nil {
		t.Fatalf("GetOptionsChains() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetOptionsChains() = %s, want %s", result, body)
	}
}

// TestGetOptionsExpirations verifies the GetOptionsExpirations method with symbol parameter.
func TestGetOptionsExpirations(t *testing.T) {
	body := `{"expirations":{"date":["2024-06-21","2024-07-19"]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/markets/options/expirations" {
			t.Errorf("path = %s, want /v1/markets/options/expirations", r.URL.Path)
		}
		if r.URL.Query().Get("symbol") != "AAPL" {
			t.Errorf("symbol = %s, want AAPL", r.URL.Query().Get("symbol"))
		}
		if r.URL.Query().Get("includeAllRoots") != "true" {
			t.Errorf("includeAllRoots = %s, want true", r.URL.Query().Get("includeAllRoots"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetOptionsExpirations("AAPL", "true", "", "", "")
	if err != nil {
		t.Fatalf("GetOptionsExpirations() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetOptionsExpirations() = %s, want %s", result, body)
	}
}

// TestGetOptionsStrikes verifies the GetOptionsStrikes method with symbol and expiration.
func TestGetOptionsStrikes(t *testing.T) {
	body := `{"strikes":{"strike":[100,105,110,115]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/markets/options/strikes" {
			t.Errorf("path = %s, want /v1/markets/options/strikes", r.URL.Path)
		}
		if r.URL.Query().Get("symbol") != "AAPL" {
			t.Errorf("symbol = %s, want AAPL", r.URL.Query().Get("symbol"))
		}
		if r.URL.Query().Get("expiration") != "2024-06-21" {
			t.Errorf("expiration = %s, want 2024-06-21", r.URL.Query().Get("expiration"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetOptionsStrikes("AAPL", "2024-06-21")
	if err != nil {
		t.Fatalf("GetOptionsStrikes() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetOptionsStrikes() = %s, want %s", result, body)
	}
}

// TestGetOptionsLookup verifies the GetOptionsLookup method with underlying and filters.
func TestGetOptionsLookup(t *testing.T) {
	body := `{"symbols":[{"rootSymbol":"AAPL","options":["AAPL220617C00270000"]}]}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/markets/options/lookup" {
			t.Errorf("path = %s, want /v1/markets/options/lookup", r.URL.Path)
		}
		if r.URL.Query().Get("underlying") != "AAPL" {
			t.Errorf("underlying = %s, want AAPL", r.URL.Query().Get("underlying"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetOptionsLookup("AAPL", "", "", "")
	if err != nil {
		t.Fatalf("GetOptionsLookup() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetOptionsLookup() = %s, want %s", result, body)
	}
}

// TestGetHistoricalPricing verifies the GetHistoricalPricing method with date range and interval.
func TestGetHistoricalPricing(t *testing.T) {
	body := `{"history":{"day":[{"date":"2024-01-02","open":150.0}]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/markets/history" {
			t.Errorf("path = %s, want /v1/markets/history", r.URL.Path)
		}
		if r.URL.Query().Get("symbol") != "AAPL" {
			t.Errorf("symbol = %s, want AAPL", r.URL.Query().Get("symbol"))
		}
		if r.URL.Query().Get("interval") != "daily" {
			t.Errorf("interval = %s, want daily", r.URL.Query().Get("interval"))
		}
		if r.URL.Query().Get("start") != "2024-01-01" {
			t.Errorf("start = %s, want 2024-01-01", r.URL.Query().Get("start"))
		}
		if r.URL.Query().Get("end") != "2024-12-31" {
			t.Errorf("end = %s, want 2024-12-31", r.URL.Query().Get("end"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetHistoricalPricing("AAPL", "daily", "2024-01-01", "2024-12-31")
	if err != nil {
		t.Fatalf("GetHistoricalPricing() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetHistoricalPricing() = %s, want %s", result, body)
	}
}

// TestGetTimeSales verifies the GetTimeSales method with interval and session filter.
func TestGetTimeSales(t *testing.T) {
	body := `{"series":{"data":[{"timestamp":"2024-01-02T10:00:00","price":150.0}]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/markets/timesales" {
			t.Errorf("path = %s, want /v1/markets/timesales", r.URL.Path)
		}
		if r.URL.Query().Get("symbol") != "AAPL" {
			t.Errorf("symbol = %s, want AAPL", r.URL.Query().Get("symbol"))
		}
		if r.URL.Query().Get("interval") != "5min" {
			t.Errorf("interval = %s, want 5min", r.URL.Query().Get("interval"))
		}
		if r.URL.Query().Get("session_filter") != "open" {
			t.Errorf("session_filter = %s, want open", r.URL.Query().Get("session_filter"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetTimeSales("AAPL", "5min", "", "", "open")
	if err != nil {
		t.Fatalf("GetTimeSales() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetTimeSales() = %s, want %s", result, body)
	}
}

// TestGetCalendar verifies the GetCalendar method with month and year parameters.
func TestGetCalendar(t *testing.T) {
	body := `{"calendar":{"month":1,"year":2024}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/markets/calendar" {
			t.Errorf("path = %s, want /v1/markets/calendar", r.URL.Path)
		}
		if r.URL.Query().Get("month") != "1" {
			t.Errorf("month = %s, want 1", r.URL.Query().Get("month"))
		}
		if r.URL.Query().Get("year") != "2024" {
			t.Errorf("year = %s, want 2024", r.URL.Query().Get("year"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetCalendar("1", "2024")
	if err != nil {
		t.Fatalf("GetCalendar() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetCalendar() = %s, want %s", result, body)
	}
}

// TestGetClock verifies the GetClock method returns market status.
func TestGetClock(t *testing.T) {
	body := `{"clock":{"date":"2024-01-02","state":"open"}}`
	server := testServer(t, "GET", "/v1/markets/clock", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetClock()
	if err != nil {
		t.Fatalf("GetClock() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetClock() = %s, want %s", result, body)
	}
}

// TestGetETB verifies the GetETB method returns easy-to-borrow securities.
func TestGetETB(t *testing.T) {
	body := `{"securities":{"security":[{"symbol":"AAPL"}]}}`
	server := testServer(t, "GET", "/v1/markets/etb", 200, body)
	defer server.Close()
	c := testClient(server)

	result, err := c.GetETB()
	if err != nil {
		t.Fatalf("GetETB() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetETB() = %s, want %s", result, body)
	}
}

// TestGetLookup verifies the GetLookup method with query and filter parameters.
func TestGetLookup(t *testing.T) {
	body := `{"securities":{"security":[{"symbol":"AAPL","exchange":"Q","type":"stock"}]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/markets/lookup" {
			t.Errorf("path = %s, want /v1/markets/lookup", r.URL.Path)
		}
		if r.URL.Query().Get("q") != "AAPL" {
			t.Errorf("q = %s, want AAPL", r.URL.Query().Get("q"))
		}
		if r.URL.Query().Get("types") != "stock" {
			t.Errorf("types = %s, want stock", r.URL.Query().Get("types"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetLookup("AAPL", "", "stock")
	if err != nil {
		t.Fatalf("GetLookup() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetLookup() = %s, want %s", result, body)
	}
}

// TestGetSearch verifies the GetSearch method with query and indexes parameters.
func TestGetSearch(t *testing.T) {
	body := `{"securities":{"security":[{"symbol":"AAPL","description":"Apple Inc"}]}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/markets/search" {
			t.Errorf("path = %s, want /v1/markets/search", r.URL.Path)
		}
		if r.URL.Query().Get("q") != "apple" {
			t.Errorf("q = %s, want apple", r.URL.Query().Get("q"))
		}
		if r.URL.Query().Get("indexes") != "true" {
			t.Errorf("indexes = %s, want true", r.URL.Query().Get("indexes"))
		}
		w.WriteHeader(200)
		w.Write([]byte(body))
	}))
	defer server.Close()
	c := testClient(server)

	result, err := c.GetSearch("apple", "true")
	if err != nil {
		t.Fatalf("GetSearch() error: %v", err)
	}
	if string(result) != body {
		t.Errorf("GetSearch() = %s, want %s", result, body)
	}
}
