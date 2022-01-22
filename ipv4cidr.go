package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-robinson/iplib"
)

func processIpv4Cidr(ipv4Cidr string) {
	net4 := iplib.Net4FromStr(ipv4Cidr)
	ip := net4.IP()
	if ip != nil {
		switch strings.ToLower(strings.TrimSpace(*flag_filter)) {
		default:
			//fmt.Printf("cidr        %v\n", ipv4Cidr)
			mask := net4.Mask()
			fmt.Printf("mask        %v.%v.%v.%v\n", mask[0], mask[1], mask[2], mask[3])
			fmt.Printf("netaddress  %v\n", ip)
			fmt.Printf("broadcast   %v\n", net4.BroadcastAddress())
			fmt.Printf("host-count  %v\n", net4.Count())
			fmt.Printf("first-host  %v\n", net4.FirstAddress())
			fmt.Printf("last-host   %v\n", net4.LastAddress())
		case "mask":
			mask := net4.Mask()
			fmt.Printf("%v.%v.%v.%v\n", mask[0], mask[1], mask[2], mask[3])
		case "netaddress":
			fmt.Println(ip)
		case "multicast", "broadcast":
			fmt.Println(net4.BroadcastAddress())
		case "host-count":
			fmt.Println(net4.Count())
		case "first-host":
			fmt.Println(net4.FirstAddress())
		case "last-host":
			fmt.Println(net4.LastAddress())
		}
	} else {
		fmt.Println("Failed to parse CIDR")
		os.Exit(1)
	}
}
