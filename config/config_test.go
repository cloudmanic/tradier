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
		APIKey:    "test-api-key-123",
		Sandbox:   true,
		AccountID: "VA000001",
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

	if loaded.APIKey != cfg.APIKey {
		t.Errorf("APIKey = %q, want %q", loaded.APIKey, cfg.APIKey)
	}
	if loaded.Sandbox != cfg.Sandbox {
		t.Errorf("Sandbox = %v, want %v", loaded.Sandbox, cfg.Sandbox)
	}
	if loaded.AccountID != cfg.AccountID {
		t.Errorf("AccountID = %q, want %q", loaded.AccountID, cfg.AccountID)
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

// TestBaseURL verifies that BaseURL returns the correct URL based on sandbox setting.
func TestBaseURL(t *testing.T) {
	cfg := &Config{Sandbox: false}
	if got := cfg.BaseURL(); got != ProductionBaseURL {
		t.Errorf("BaseURL() = %q, want %q", got, ProductionBaseURL)
	}

	cfg.Sandbox = true
	if got := cfg.BaseURL(); got != SandboxBaseURL {
		t.Errorf("BaseURL() = %q, want %q", got, SandboxBaseURL)
	}
}
