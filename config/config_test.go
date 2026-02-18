// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package config

import (
	"os"
	"path/filepath"
	"testing"
)

// TestSaveAndLoad verifies that saving and loading a config round-trips correctly.
func TestSaveAndLoad(t *testing.T) {
	// Create a temp directory to act as home
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	cfg := &Config{
		ProductionAPIKey:    "prod-key-123",
		ProductionAccountID: "VA000001",
		SandboxAPIKey:       "sandbox-key-456",
		SandboxAccountID:    "VA000002",
	}

	// Save the config
	err := Save(cfg)
	if err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	// Verify the file was created with correct permissions
	path := filepath.Join(tmpDir, ".config", "tradier", "config.json")
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("config file not created: %v", err)
	}
	if info.Mode().Perm() != 0600 {
		t.Errorf("config file permissions = %v, want 0600", info.Mode().Perm())
	}

	// Load the config back
	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if loaded.ProductionAPIKey != cfg.ProductionAPIKey {
		t.Errorf("ProductionAPIKey = %q, want %q", loaded.ProductionAPIKey, cfg.ProductionAPIKey)
	}
	if loaded.ProductionAccountID != cfg.ProductionAccountID {
		t.Errorf("ProductionAccountID = %q, want %q", loaded.ProductionAccountID, cfg.ProductionAccountID)
	}
	if loaded.SandboxAPIKey != cfg.SandboxAPIKey {
		t.Errorf("SandboxAPIKey = %q, want %q", loaded.SandboxAPIKey, cfg.SandboxAPIKey)
	}
	if loaded.SandboxAccountID != cfg.SandboxAccountID {
		t.Errorf("SandboxAccountID = %q, want %q", loaded.SandboxAccountID, cfg.SandboxAccountID)
	}
}

// TestLoadMissingFile verifies that Load returns an error when no config exists.
func TestLoadMissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	_, err := Load()
	if err == nil {
		t.Fatal("Load() should return error when config file is missing")
	}
}

// TestBaseURL verifies that BaseURL returns the correct URL based on sandbox flag.
func TestBaseURL(t *testing.T) {
	cfg := &Config{}

	if got := cfg.BaseURL(false); got != ProductionBaseURL {
		t.Errorf("BaseURL(false) = %q, want %q", got, ProductionBaseURL)
	}

	if got := cfg.BaseURL(true); got != SandboxBaseURL {
		t.Errorf("BaseURL(true) = %q, want %q", got, SandboxBaseURL)
	}
}

// TestAPIKey verifies that APIKey returns the correct key based on sandbox flag.
func TestAPIKey(t *testing.T) {
	cfg := &Config{
		ProductionAPIKey: "prod-key",
		SandboxAPIKey:    "sandbox-key",
	}

	if got := cfg.APIKey(false); got != "prod-key" {
		t.Errorf("APIKey(false) = %q, want %q", got, "prod-key")
	}

	if got := cfg.APIKey(true); got != "sandbox-key" {
		t.Errorf("APIKey(true) = %q, want %q", got, "sandbox-key")
	}
}

// TestAccountID verifies that AccountID returns the correct ID based on sandbox flag.
func TestAccountID(t *testing.T) {
	cfg := &Config{
		ProductionAccountID: "prod-acct",
		SandboxAccountID:    "sandbox-acct",
	}

	if got := cfg.AccountID(false); got != "prod-acct" {
		t.Errorf("AccountID(false) = %q, want %q", got, "prod-acct")
	}

	if got := cfg.AccountID(true); got != "sandbox-acct" {
		t.Errorf("AccountID(true) = %q, want %q", got, "sandbox-acct")
	}
}
