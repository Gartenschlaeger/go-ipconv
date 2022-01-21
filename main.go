package main

import (
	"os"
	"regexp"
	"strings"
)

// ipv4 address : 000.000.000.000
const ip4vPattern = "^\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3}$"

// ipv4 cidr : 000.000.000.000/00
const ipv4CidrPattern = "^\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3}/\\d{1,2}$"

func main() {
	argsCount := len(os.Args)
	if argsCount < 2 {
		writeHelpOutput()
	} else {
		firstArgument := strings.TrimSpace(os.Args[1])

		ipv4Address, _ := regexp.MatchString(ip4vPattern, firstArgument)
		if ipv4Address {
			processIpv4Address(firstArgument)
			return
		}

		cidrMatch, _ := regexp.MatchString(ipv4CidrPattern, firstArgument)
		if cidrMatch {
			processIpv4Cidr(firstArgument)
			return
		}

		writeHelpOutput()
	}
}
