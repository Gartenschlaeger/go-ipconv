package main

import (
	"os"
	"regexp"
	"strings"
)

func main() {
	argsCount := len(os.Args)
	if argsCount < 2 {
		writeHelpOutput()
	} else {
		firstArgument := strings.TrimSpace(os.Args[1])

		// handle ipv4 address
		ipv4Address, _ := regexp.MatchString(ip4vPattern, firstArgument)
		if ipv4Address {
			processIpv4Address(firstArgument)
			return
		}

		// handle ipv4 cidr
		cidrMatch, _ := regexp.MatchString(ipv4CidrPattern, firstArgument)
		if cidrMatch {
			processIpv4Cidr(firstArgument)
			return
		}

		writeHelpOutput()
	}
}
