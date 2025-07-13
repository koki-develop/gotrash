# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Build
```bash
go build .
```

### Lint
```bash
golangci-lint run ./...
```

### Release Check
```bash
goreleaser check
```

### Run the Application
```bash
go run . [command]
# or after building:
./gotrash [command]
```

## Architecture Overview

gotrash is a trash/recycle bin implementation in Go that provides safer file deletion by moving files to a trash directory instead of permanently deleting them.

### Core Components

1. **Command Structure** (cmd/)
   - `root.go`: Main command setup and flag definitions
   - `put.go`: Moves files/directories to trash
   - `list.go`: Lists trashed items
   - `restore.go`: Restores files from trash (supports interactive fuzzy finder)
   - `clear.go`: Permanently deletes all trashed items

2. **Database Layer** (internal/db/db.go)
   - Uses BuntDB for tracking trashed items
   - Trash location: `$GOTRASH_ROOT/can` (default: `$HOME/.gotrash/can`)
   - Database file: `$GOTRASH_ROOT/db`
   - Auto-shrinks database when it exceeds 10MB

3. **Data Model** (internal/trash/trash.go)
   - Trash items are stored with unique keys: `{unix_timestamp}_{uuid}`
   - Tracks original path and trash timestamp
   - Implements fuzzy finder interface

4. **Utilities** (internal/util/)
   - File operations helpers
   - Interactive prompts and filters

### Key Design Decisions

- Files are moved (not copied) to trash directory with unique identifiers
- Original paths are preserved in database for restoration
- Interactive restore uses go-fzf for fuzzy finding
- Transactional updates ensure database and filesystem stay in sync