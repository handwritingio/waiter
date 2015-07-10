package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"time"
)

const timeout = 5 * time.Second

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}

	parsedUrl, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	address := parsedUrl.Host
	if address == "" {
		address = os.Args[1]
	}

	fmt.Fprintf(os.Stdout, "Waiting for TCP connection to %s...\n", address)
	if _, err := net.DialTimeout("tcp", address, timeout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
