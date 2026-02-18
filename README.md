# Tradier CLI

A command-line interface for the [Tradier](https://tradier.com) brokerage API. Manage your accounts, query real-time market data, place trades, and more -- all from your terminal.

## Why a CLI?

**A context-friendly tool for AI agents.** MCP servers can flood an AI agent's context window with large JSON payloads, eating up tokens and degrading response quality. A CLI tool lets your agent run targeted queries and pipe just the relevant output into the conversation. Smaller, more focused context means better results.

**Better for debugging.** When something breaks in an API integration, you can run the exact same command yourself, see the raw output, tweak flags, and figure out what went wrong in seconds.

**Works everywhere.** A single binary. No SDK to install, no server to run. It works in your terminal, in shell scripts, in CI pipelines, and as a tool that any AI coding assistant can shell out to.

**Human-readable by default, machine-readable when you need it.** Every command outputs clean, styled tables for scanning by eye, and switches to JSON with a single `--json` flag for piping into `jq`, scripts, or AI context.

## Installation

### Homebrew (macOS / Linux)

```bash
brew tap cloudmanic/tradier https://github.com/cloudmanic/tradier
brew install tradier
```

To upgrade to the latest version:

```bash
brew upgrade tradier
```

### Pre-built Binaries

Download the latest binary for your platform from the [Releases](https://github.com/cloudmanic/tradier/releases) page.

Available platforms: macOS (Intel & Apple Silicon), Linux (amd64 & arm64), Windows (amd64).

### From Source

```bash
git clone https://github.com/cloudmanic/tradier.git
cd tradier
make build
```

The binary will be at `build/tradier`. Move it somewhere in your `$PATH`:

```bash
mv build/tradier /usr/local/bin/
```

### Requirements

- A [Tradier](https://tradier.com) brokerage or sandbox account with an API key
- Go 1.24+ (only if building from source)

## Quick Start

### 1. Configure your API key

```bash
tradier init
```

This prompts for your API key and optional default account ID, then saves the configuration to `~/.config/tradier/config.json`.

For sandbox accounts, use the `--sandbox` flag:

```bash
tradier init --sandbox
```

### 2. Start querying

```bash
# Get a real-time quote
tradier markets quotes --symbols AAPL,MSFT

# Check your account balance
tradier accounts balance

# View recent orders
tradier accounts orders

# Get historical prices
tradier markets history --symbol AAPL --start 2025-01-01 --end 2025-12-31

# Everything outputs JSON too
tradier markets quotes --symbols AAPL --json
```

## Configuration

Configuration is stored at `~/.config/tradier/config.json` with `0600` permissions (owner-only access).

| Field | Description |
|-------|-------------|
| `api_key` | Your Tradier API access token |
| `sandbox` | `true` for sandbox environment, `false` for production |
| `account_id` | Default account ID (avoids needing `--account-id` on every command) |

- **Production URL:** `https://api.tradier.com`
- **Sandbox URL:** `https://sandbox.tradier.com`

## Output Formats

Every command supports two output formats:

```bash
# Styled table output (default) -- human-readable, aligned columns
tradier accounts history

# JSON output -- machine-readable, pipe to jq or feed to an AI agent
tradier accounts history --json
```

Option symbols are automatically displayed in human-readable format:

| OCC Format | Display |
|------------|---------|
| `SPY260220P00657000` | `SPY 02/20/26 $657 Put` |
| `AAPL220617C00270000` | `AAPL 06/17/22 $270 Call` |

## Commands

### Account Management

```bash
# Account balance and margin information
tradier accounts balance

# Current positions
tradier accounts positions

# All orders (with multileg indicator)
tradier accounts orders

# Specific order detail (shows all legs)
tradier accounts order --order-id 12345

# Closed position gain/loss
tradier accounts gainloss
tradier accounts gainloss --sort-by gainloss --sort desc --limit 20

# Account history (trades, ACH, fees, dividends, etc.)
tradier accounts history
tradier accounts history --type trade --start 2025-01-01 --end 2025-12-31

# Historical balance over time
tradier accounts historical-balances --period MONTH
```

### Position Groups

```bash
# List all position groups
tradier accounts position-groups list

# Create a new group
tradier accounts position-groups create --label "Tech Stocks" --symbols "AAPL,MSFT,GOOGL"

# Update a group
tradier accounts position-groups update --group-id abc123 --label "Big Tech" --symbols "AAPL,MSFT,GOOGL,AMZN"

# Delete a group
tradier accounts position-groups delete --group-id abc123
```

### Market Data

```bash
# Real-time quotes
tradier markets quotes --symbols AAPL,MSFT,GOOGL

# Quotes via POST (for large symbol lists)
tradier markets post-quotes --symbols "AAPL,MSFT,GOOGL,AMZN,META,NFLX,TSLA"

# Historical OHLCV pricing
tradier markets history --symbol AAPL
tradier markets history --symbol AAPL --interval weekly --start 2025-01-01 --end 2025-12-31

# Time and sales (intraday)
tradier markets timesales --symbol AAPL --interval 5min --start "2025-06-15 09:30" --end "2025-06-15 16:00"

# Market calendar
tradier markets calendar --month 3 --year 2025

# Market clock (current status)
tradier markets clock

# Easy-to-borrow list
tradier markets etb

# Symbol lookup
tradier markets lookup --query AAPL
tradier markets search --query "Apple"
```

### Options

```bash
# Option chains for a specific expiration
tradier markets options-chains --symbol AAPL --expiration 2025-06-20

# Available expiration dates
tradier markets options-expirations --symbol AAPL

# Available strike prices
tradier markets options-strikes --symbol AAPL --expiration 2025-06-20

# Look up option symbols
tradier markets options-lookup --underlying AAPL
tradier markets options-lookup --underlying AAPL --type call --strike 200 --expiration 2025-06-20
```

### Trading

```bash
# Equity market buy
tradier trading place --class equity --symbol AAPL --side buy --quantity 10 --type market --duration day

# Equity limit sell
tradier trading place --class equity --symbol AAPL --side sell --quantity 10 --type limit --duration day --price 200.00

# Option buy to open
tradier trading place --class option --symbol AAPL --option-symbol AAPL220617C00270000 \
  --side buy_to_open --quantity 5 --type limit --duration day --price 3.50

# Multileg spread
tradier trading place --class multileg --symbol AAPL --type debit --duration day --price 1.50 \
  --option-symbol-0 AAPL220617C00270000 --side-0 buy_to_open --quantity-0 1 \
  --option-symbol-1 AAPL220617C00280000 --side-1 sell_to_open --quantity-1 1

# Preview an order (validates without submitting)
tradier trading place --class equity --symbol AAPL --side buy --quantity 10 --type market --duration day --preview true

# Modify an existing order
tradier trading change --order-id 12345 --type limit --price 205.00

# Cancel an order
tradier trading cancel --order-id 12345
```

**Supported order classes:** `equity`, `option`, `multileg`, `combo`, `oto`, `oco`, `otoco`

**Supported sides:** `buy`, `sell`, `sell_short`, `buy_to_cover`, `buy_to_open`, `buy_to_close`, `sell_to_open`, `sell_to_close`

**Supported order types:** `market`, `limit`, `stop`, `stop_limit`, `debit`, `credit`, `even`

**Supported durations:** `day`, `gtc`, `pre`, `post`

### Watchlists

```bash
# List all watchlists
tradier watchlists list

# Get a specific watchlist
tradier watchlists get --id my-watchlist-id

# Create a new watchlist
tradier watchlists create --name "My Watchlist" --symbols "AAPL,MSFT,GOOGL"

# Update a watchlist
tradier watchlists update --id my-watchlist-id --name "Updated Name" --symbols "AAPL,TSLA"

# Add symbols to a watchlist
tradier watchlists add-symbols --id my-watchlist-id --symbols "AMZN,META"

# Remove a symbol from a watchlist
tradier watchlists remove-symbol --id my-watchlist-id --symbol META

# Delete a watchlist
tradier watchlists delete --id my-watchlist-id
```

### User Profile

```bash
# Get user profile and linked accounts
tradier user profile
```

### Streaming Sessions

```bash
# Create a market data streaming session
tradier streaming market-session

# Create an account events streaming session
tradier streaming account-session
```

## Using with AI Agents

Tradier CLI is designed to work well as a tool for AI coding assistants. Any agent that can execute shell commands can use it.

### Pipe JSON into context

```bash
# Get structured data an agent can reason about
tradier accounts orders --json | head -50

# Check positions for analysis
tradier accounts positions --json

# Get option chain data
tradier markets options-chains --symbol SPY --expiration 2025-06-20 --json
```

### Use in scripts

```bash
#!/bin/bash
# Monitor multiple positions
for symbol in AAPL MSFT GOOGL AMZN; do
  echo "=== $symbol ==="
  tradier markets quotes --symbols "$symbol"
  echo
done
```

### Check market status before trading

```bash
# Is the market open?
tradier markets clock --json | jq '.clock.state'
```

## Global Flags

| Flag | Description |
|------|-------------|
| `--json` | Output raw JSON instead of formatted tables |
| `--account-id` | Override the default account ID from config (on account/trading commands) |
| `--version` | Print the CLI version |
| `--help` | Show help for any command |

## Development

```bash
# Build
make build

# Run all tests
make test

# Run tests with verbose output
make test-verbose

# Test coverage
make test-cover

# Format and lint
make lint

# Cross-compile for all platforms
make cross-build

# See all targets
make help
```

## License

Copyright (c) 2026 Cloudmanic Labs, LLC. All rights reserved.
