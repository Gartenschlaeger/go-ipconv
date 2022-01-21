package main

import (
	"fmt"
	"net"

	"github.com/c-robinson/iplib"
)

func processIpv4Address(ipv4Address string) {
	ip := net.ParseIP(ipv4Address)
	if ip != nil {
		fmt.Printf("address         %v\n", ipv4Address)
		fmt.Printf("arpa            %v\n", iplib.IP4ToARPA(ip))
		fmt.Printf("is private      %v\n", ip.IsPrivate())
		fmt.Printf("is loopback     %v\n", ip.IsLoopback())
		fmt.Printf("is multicast    %v\n", ip.IsMulticast())
		fmt.Printf("is unspecified  %v\n", ip.IsUnspecified())
		fmt.Printf("dec             %v\n", iplib.IP4ToUint32(ip))
		fmt.Printf("bin             %v\n", iplib.IPToBinaryString(ip))
		fmt.Printf("hex             %v\n", iplib.IPToHexString(ip))
	}
}
