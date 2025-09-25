# Auto DNS Query

A simple command-line tool that continuously queries DNS hosts and tracks unique IP addresses.

## Features

- Query DNS hosts at regular intervals
- Bypass OS DNS cache for fresh results
- Filter by IP type (IPv4, IPv6, or both)
- Track and save unique IP addresses to file
- Cross-platform support (Linux, macOS, Windows)
- Built-in deduplication
- Configurable query intervals and output files

## Usage

```bash
./autoDnsQuery -host example.com -interval 10 -output results.txt -mode both
```

### Options

- `-host` - DNS hostname to query (required)
- `-interval` - Query interval in seconds (default: 5)
- `-output` - Output file for IP addresses (default: result.txt)
- `-mode` - IP mode: "ipv4", "ipv6", or "both" (default: ipv4)

## Building

### Quick build
```bash
make build
```

### Cross-platform build
```bash
make build-all  # Builds for Linux, Windows, and macOS
```

### Individual platforms
```bash
make linux      # Linux binary
make windows    # Windows binary
make darwin     # macOS binary
```

## Example

```bash
$ ./bin/autoDnsQuery -host google.com -interval 5 -mode both
Querying google.com every 5 seconds...
 New IP found: 142.250.191.14 (total: 1)
 New IP found: 2607:f8b0:4004:c1b::65 (total: 2)
IP already recorded: 142.250.191.14
```

The tool bypasses OS DNS cache and will continuously monitor the specified host, saving all unique IP addresses to the output file with fresh DNS lookups.