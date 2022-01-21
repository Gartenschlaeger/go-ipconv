package main

import (
	"fmt"
	"os"

	"github.com/c-robinson/iplib"
)

func processIpv4Cidr(ipv4Cidr string) {
	net4 := iplib.Net4FromStr(ipv4Cidr)
	ip := net4.IP()

	if ip == nil {
		fmt.Println("Failed to parse CIDR")
		os.Exit(1)
		return
	}

	mask := net4.Mask()

	fmt.Printf("cidr           %v\n", ipv4Cidr)
	fmt.Printf("mask           %v.%v.%v.%v\n", mask[0], mask[1], mask[2], mask[3])
	fmt.Printf("net            %v\n", ip)
	fmt.Printf("broadcast      %v\n", net4.BroadcastAddress())
	fmt.Printf("host count     %v\n", net4.Count())
	fmt.Printf("first host     %v\n", net4.FirstAddress())
	fmt.Printf("last host      %v\n", net4.LastAddress())
}
