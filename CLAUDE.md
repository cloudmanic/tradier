# CLAUDE.md

## Project Overview

Tradier CLI is a Go command-line tool for the [Tradier brokerage API](https://tradier.com). It covers all 37 API endpoints across accounts, market data, trading, watchlists, streaming, and user management.

**API Reference:** https://docs.tradier.com/llms.txt

## Architecture

```
main.go              Entry point, calls cmd.Execute()
config/              Config loading/saving (~/.config/tradier/config.json)
client/              HTTP client + all API methods (one file per domain)
cmd/                 Cobra CLI commands (one file per domain) + display.go
Formula/             Homebrew formula
.github/workflows/   CI (test on PR) and CD (release on merge to main)
```

### Package Responsibilities

- **config/** - Stores both production and sandbox API keys/account IDs. The `--sandbox` flag on any command selects which set of credentials to use.
- **client/** - Stateless HTTP client. All methods return raw `[]byte` JSON. Uses `doGet`, `doPost`, `doPut`, `doDelete` helpers. No business logic.
- **cmd/** - Cobra command tree. Each command calls a client method and passes the result to `printResult(data, displayFunc)`.
- **cmd/display.go** - All table rendering logic. JSON navigation helpers (`parseJSON`, `nested`, `toSlice`, `str`, `num`), formatting helpers (`money`, `pct`, `formatOptionSymbol`), and ~25 display functions.

### Key Patterns

- **Output:** Default is styled tables via `go-pretty` with `StyleRounded`. Pass `--json` for raw JSON.
- **OCC Option Symbols:** Always display in human-readable format (`SPY 02/20/26 $657 Put`) using `formatOptionSymbol()` in display.go.
- **Tradier API Quirk:** Single items come back as JSON objects, multiple as arrays. The `toSlice()` helper normalizes both to `[]map[string]interface{}`.
- **Account type sub-objects:** Balance data nests buying power under the account type key (`pdt`, `margin`, or `cash`). The display function detects this dynamically.
- **Multileg orders:** In the orders table, show first leg's option symbol with `*` suffix and a legend at the bottom.

## Build and Test

```bash
make build          # Build to build/tradier with version ldflags
make test           # Run all tests
make test-verbose   # Verbose test output
make test-cover     # Coverage report
make cross-build    # Build for all platforms (dist/)
make lint           # Format + vet
```

Version is injected at build time via `-ldflags "-X github.com/cloudmanic/tradier/cmd.version=..."`.

## Testing Conventions

- One test file per source file (e.g., `accounts.go` -> `accounts_test.go`)
- All API calls are mocked using `httptest.NewServer`
- Test helpers in `client/client_test.go`: `testServer()` and `testClient()`
- Mock data must use obviously fake values (e.g., `VA000001`, `John Doe`, `test-key`). Never use real account data.
- Tests verify HTTP method, path, query parameters, and response parsing

## Code Style

- Copyright header on every new file: `// Copyright 2026 Cloudmanic Labs, LLC. All rights reserved.` with `// Date: YYYY-MM-DD`
- Detailed comments above every function, public and private
- No third-party test frameworks -- standard `testing` package only
- Errors bubble up via `RunE` in cobra commands, no `os.Exit` in commands
- All API client methods return `([]byte, error)` -- raw JSON bytes, no struct unmarshaling

## CI/CD

- **PRs:** `.github/workflows/test.yml` runs `go test ./...` on every PR to main
- **Releases:** `.github/workflows/release.yml` on push to main: runs tests, auto-increments semver (patch), builds cross-platform binaries, creates GitHub release with tag
- **Homebrew:** `Formula/tradier.rb` points to latest release binaries. Install via `brew tap cloudmanic/tradier https://github.com/cloudmanic/tradier`

## Config File

Located at `~/.config/tradier/config.json` with `0600` permissions:

```json
{
  "production_api_key": "...",
  "production_account_id": "...",
  "sandbox_api_key": "...",
  "sandbox_account_id": "..."
}
```

## Adding a New Endpoint

1. Add the API method to the appropriate file in `client/` (returns `[]byte, error`)
2. Add a test in the corresponding `_test.go` using `testServer()` mock
3. Add a display function in `cmd/display.go`
4. Add a cobra command in the appropriate `cmd/` file, wire it into `init()`
5. Run `make test` and `make build`
