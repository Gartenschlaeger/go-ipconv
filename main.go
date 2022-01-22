package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var regexp_ipv4 = regexp.MustCompile(`(?m)^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
var regexp_ipv4cidr = regexp.MustCompile(`(?m)^(\d{1,3})(\.\d{1,3})(\.\d{1,3})?(\.\d{1,3})?(\/\d{1,2})$`)

var flag_input *string
var flag_filter *string

func main() {
	flag_input = flag.String("i", "", "address pattern (required)")
	flag_filter = flag.String("f", "", "[dec, bin, hex]")
	flag.Parse()

	if strings.TrimSpace(*flag_input) == "" {
		fmt.Println("ipconv -i <address pattern>")
		return
	}

	addressPattern := strings.TrimSpace(*flag_input)

	// ipv4 ip address
	ipv4Address := regexp_ipv4.MatchString(addressPattern)
	if ipv4Address {
		processIpv4Address()
		return
	}

	// ipv4 cidr
	cidrMatch := regexp_ipv4cidr.MatchString(addressPattern)
	if cidrMatch {
		processIpv4Cidr(addressPattern)
		return
	}

	fmt.Println("Input pattern has an invalid format. ipv4 address or CIDR expected.")
	os.Exit(1)
}
