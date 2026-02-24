# sync-once-2

This package provides an alternative or extended example of using Go's `sync.Once` primitive.

## Overview

Like the first `sync-once` package, this demonstrates how to ensure a function is executed only once, even in the presence of multiple goroutines. You can use this as a reference for different patterns or variations on one-time initialization.

## Usage

- Run the example with:
  ```sh
  go run main.go
  ```
- Observe that the initialization logic is executed only once, regardless of concurrent calls.
