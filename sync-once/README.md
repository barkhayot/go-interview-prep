# sync-once

This package demonstrates the use of Go's `sync.Once` primitive.

## Overview

The `sync.Once` type ensures that a function is only executed once, even if called from multiple goroutines. This is useful for one-time initialization logic in concurrent programs.

## Usage

- Run the example with:
  ```sh
  go run main.go
  ```
- The output will show that the initialization logic is executed only once, regardless of how many goroutines attempt to run it.
