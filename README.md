# gotrash

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/koki-develop/gotrash)](https://github.com/koki-develop/gotrash/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/koki-develop/gotrash/ci.yml?logo=github)](https://github.com/koki-develop/gotrash/actions/workflows/ci.yml)
[![Maintainability](https://img.shields.io/codeclimate/maintainability/koki-develop/gotrash?style=flat&logo=codeclimate)](https://codeclimate.com/github/koki-develop/gotrash/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/koki-develop/gotrash)](https://goreportcard.com/report/github.com/koki-develop/gotrash)
[![LICENSE](https://img.shields.io/github/license/koki-develop/gotrash)](./LICENSE)

rm alternative written in Go.

- [Installation](#installation)
- [Usage](#usage)
  - [`put`](#gotrash-put)
  - [`list`](#gotrash-list)
  - [`restore`](#gotrash-restore)
  - [`clear`](#gotrash-clear)
- [LICENSE](#license)

## Installation

### Homebrew

```console
$ brew install koki-develop/tap/gotrash
```

### `go install`

```console
$ go install github.com/koki-develop/gotrash@latest
```

### Releases

Download the binary from the [releases page](https://github.com/koki-develop/gotrash/releases/latest).

## Usage

```console
$ gotrash --help
rm alternative written in Go.

Usage:
  gotrash [command]

Available Commands:
  clear       Clear all trashed files or directories
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List trashed flies or directories
  put         Trash files or directories
  restore     Restore trashed files or directories

Flags:
  -h, --help      help for gotrash
  -v, --version   version for gotrash

Use "gotrash [command] --help" for more information about a command.
```

### `gotrash put`

```console
$ gotrash put --help
Trash files or directories.

Usage:
  gotrash put [file]... [flags]

Flags:
  -h, --help   help for put
```

### `gotrash list`

```console
$ gotrash list --help
List trashed flies or directories.

Usage:
  gotrash list [flags]

Aliases:
  list, ls

Flags:
  -c, --current-dir   show only the trash in the current directory
  -h, --help          help for list
```

### `gotrash restore`

```console
$ gotrash restore --help
Restore trashed files or directories.

Usage:
  gotrash restore [index]... [flags]

Aliases:
  restore, rs

Flags:
  -f, --force   overwrite a file or directory if it already exists
  -h, --help    help for restore
```

### `gotrash clear`

```console
$ gotrash clear --help
Clear all trashed files or directories.

Usage:
  gotrash clear [flags]

Flags:
  -f, --force   skip confirmation before clear
  -h, --help    help for clear
```

## LICENSE

[MIT](./LICENSE)
