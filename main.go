package main

import (
	"log"
	"net"
	"net/url"
	"os"
	"time"
)

const timeout = 5 * time.Second

func main() {
	if len(os.Args) < 2 {
		log.Printf("Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}

	parsedUrl, err := url.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	address := parsedUrl.Host
	if address == "" {
		address = os.Args[1]
	}

	log.Printf("Waiting for TCP connection to %s...\n", address)

	_, err = net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("ok")
}
