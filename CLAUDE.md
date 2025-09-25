# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a simple Go command-line tool that performs DNS queries at regular intervals and tracks unique IP addresses. The application uses Go's built-in `net` package for DNS resolution and saves discovered IPs to a file with deduplication.

## Architecture

The application consists of a single `main.go` file with these key components:
- Command-line argument parsing using Go's `flag` package
- DNS resolution using `net.LookupIP()`
- IP deduplication using an in-memory map
- File I/O for persisting and loading IP addresses

## Common Commands

### Building
```bash
make build              # Build for current platform to bin/
make build-all          # Build for all platforms (Linux, Windows, macOS)
make linux              # Build for Linux
make windows            # Build for Windows  
make darwin             # Build for macOS
```

### Running
```bash
make run                # Build and run with default args
./bin/autoDnsQuery -host example.com -interval 10 -output results.txt
```

### Cleaning
```bash
make clean              # Remove bin/ directory
```

### Testing  
```bash
make test               # Run Go tests
```

## Application Usage

The tool accepts these command-line arguments:
- `-host` (required): DNS hostname to query
- `-interval` (default: 5): Query interval in seconds  
- `-output` (default: "result.txt"): Output file for IP addresses

The application continuously queries the specified host, tracks unique IP addresses in memory, and appends new IPs to the output file. It uses `log.Printf` for all output with emoji indicators (✅ for new IPs, ❌ for errors).

## Code Style Notes

- Uses `log.Printf` consistently for all output (not `fmt.Printf`)
- Functions accept filename parameters rather than using global constants
- IP deduplication is handled in-memory with a `map[string]bool`