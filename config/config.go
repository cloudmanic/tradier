// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.
// Date: 2026-02-17

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// ProductionBaseURL is the base URL for the Tradier production API.
	ProductionBaseURL = "https://api.tradier.com"

	// SandboxBaseURL is the base URL for the Tradier sandbox API.
	SandboxBaseURL = "https://sandbox.tradier.com"

	// configDir is the directory name under ~/.config for storing tradier configuration.
	configDir = "tradier"

	// configFile is the configuration file name.
	configFile = "config.json"
)

// Config holds the configuration for the Tradier CLI tool.
type Config struct {
	APIKey    string `json:"api_key"`
	Sandbox   bool   `json:"sandbox"`
	AccountID string `json:"account_id"`
}

// ConfigDirPath returns the full path to the tradier configuration directory.
func ConfigDirPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to determine home directory: %w", err)
	}
	return filepath.Join(home, ".config", configDir), nil
}

// ConfigFilePath returns the full path to the tradier configuration file.
func ConfigFilePath() (string, error) {
	dir, err := ConfigDirPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, configFile), nil
}

// Load reads the configuration from disk. Returns an error if the file does not exist or cannot be parsed.
func Load() (*Config, error) {
	path, err := ConfigFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file (run 'tradier init' first): %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unable to parse config file: %w", err)
	}

	return &cfg, nil
}

// Save writes the configuration to disk, creating the config directory if needed.
func Save(cfg *Config) error {
	dir, err := ConfigDirPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("unable to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal config: %w", err)
	}

	path := filepath.Join(dir, configFile)
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("unable to write config file: %w", err)
	}

	return nil
}

// BaseURL returns the appropriate base URL based on the sandbox setting.
func (c *Config) BaseURL() string {
	if c.Sandbox {
		return SandboxBaseURL
	}
	return ProductionBaseURL
}
