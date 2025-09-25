package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	var host string
	var interval int
	var outputFile string
	var mode string

	flag.StringVar(&host, "host", "", "DNS host to query")
	flag.IntVar(&interval, "interval", 5, "Query interval in seconds")
	flag.StringVar(&outputFile, "output", "result.txt", "Output file for IP addresses")
	flag.StringVar(&mode, "mode", "ipv4", "IP mode: ipv4, ipv6, or both")
	flag.Parse()

	if host == "" {
		log.Fatal("Host is required. Use -host flag")
	}

	if mode != "ipv4" && mode != "ipv6" && mode != "both" {
		log.Fatal("Mode must be ipv4, ipv6, or both")
	}

	log.Printf("Querying %s every %d seconds...\n", host, interval)

	existingIPs := loadExistingIPs(outputFile)

	for {
		ips, err := net.LookupIP(host)
		if err != nil {
			log.Printf("❌ DNS lookup failed: %v", err)
		} else {
			for _, ip := range ips {
				ipStr := ip.String()
				
				// Filter based on mode
				isIPv4 := ip.To4() != nil
				if mode == "ipv4" && !isIPv4 {
					continue
				}
				if mode == "ipv6" && isIPv4 {
					continue
				}
				
				if !existingIPs[ipStr] {
					existingIPs[ipStr] = true
					if err := appendIPToFile(ipStr, outputFile); err != nil {
						log.Printf("❌ Failed to write IP to file: %v", err)
					} else {
						log.Printf("✅ New IP found: %s (total: %d)\n", ipStr, len(existingIPs))
					}
				} else {
					log.Printf("IP already recorded: %s\n", ipStr)
				}
			}
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func loadExistingIPs(filename string) map[string]bool {
	ips := make(map[string]bool)

	file, err := os.Open(filename)
	if err != nil {
		return ips
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if ip != "" {
			ips[ip] = true
		}
	}

	return ips
}

func appendIPToFile(ip, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(ip + "\n")
	return err
}
