package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const timeout = time.Second * 2

var port string

func main() {
	args := os.Args[1:]
	host := args[0]
	port = args[1]

	if len(args) == 0 || len(args) < 2 {
		fmt.Println("Usage: digport <host> <port>")
		os.Exit(0)
	}

	fmt.Println("Timeout is set to ", timeout, " seconds.")

	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}

	ok := 0
	nok := 0

	for _, ip := range ips {
		// skip if not ipv4
		if ip.To4() == nil {
			continue
		}

		ipstring := ip.String()

		if isPortOpen(ipstring) {
			fmt.Println(ipstring, " OK")
			ok++
			continue
		}
		nok++
		fmt.Println(ipstring + " X")
	}
	fmt.Println()
	fmt.Println("OK: " + strconv.Itoa(ok))
	fmt.Println("NOK: " + strconv.Itoa(nok))
}

func isPortOpen(ip string) bool {
	uri := ip + ":" + port
	_, err := net.DialTimeout("tcp", uri, timeout)
	if err != nil {
		return false
	}
	return true
}
