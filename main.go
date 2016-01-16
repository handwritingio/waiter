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

	timeLeft := timeout
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
