package main

import (
	"os"
	"strings"
)

func main() {
	argsCount := len(os.Args)
	if argsCount < 2 {
		writeHelpOutput()
	} else {
		firstArgument := strings.TrimSpace(os.Args[1])

		// handle ipv4 address
		ipv4Address := regexp_ipv4.MatchString(firstArgument)
		if ipv4Address {
			processIpv4Address(firstArgument)
			return
		}

		// handle ipv4 cidr
		cidrMatch := regexp_ipv4cidr.MatchString(firstArgument)
		if cidrMatch {
			processIpv4Cidr(firstArgument)
			return
		}

		writeHelpOutput()
	}
}
