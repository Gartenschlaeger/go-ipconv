package main

import "fmt"

func processIpv4Address(ipv4Address string) {
	octets := ipv4StringToOctets(ipv4Address)
	decimal := ip4vOctetsToDecimal(octets)

	fmt.Printf("IPv4 address %v\n", ipv4Address)
	fmt.Printf("Decimal      %v\n", decimal)
}
