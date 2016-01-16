package main

import (
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"time"
)

var timeoutSeconds int

func init() {
	const (
		defaultTimeout = 5 // seconds
		timeoutUsage   = "timeout in seconds"
	)
	flag.IntVar(&timeoutSeconds, "timeout", defaultTimeout, timeoutUsage)
	flag.IntVar(&timeoutSeconds, "t", defaultTimeout, timeoutUsage+" (shorthand)")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s [flags] URL\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// 1 required positional argument
	if len(flag.Args()) != 1 {
		usage()
		os.Exit(2)
	}

	rawURL := flag.Args()[0]

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	address := parsedURL.Host
	if address == "" {
		address = rawURL
	}

	timeLeft := time.Duration(timeoutSeconds) * time.Second
	for timeLeft > 0 {
		start := time.Now()
		_, err = net.DialTimeout("tcp", address, timeLeft)
		if err == nil {
			break
		}
		// DialTimeout will error out immediately if it gets "connection
		// refused" but we want to retry in that case.
		time.Sleep(100 * time.Millisecond)
		timeLeft -= time.Now().Sub(start)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
