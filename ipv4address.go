package main

import (
	"fmt"
	"net"

	"github.com/c-robinson/iplib"
)

func processIpv4Address(ipv4Address string) {
	ip := net.ParseIP(ipv4Address)
	if ip != nil {
		fmt.Printf("address        %v\n", ipv4Address)
		fmt.Printf("decimal        %v\n", iplib.IP4ToUint32(ip))
		fmt.Printf("arpa           %v\n", iplib.IP4ToARPA(ip))
		fmt.Printf("private        %v\n", ip.IsPrivate())
		fmt.Printf("loopback       %v\n", ip.IsLoopback())
		fmt.Printf("multicast      %v\n", ip.IsMulticast())
		fmt.Printf("unspecified    %v\n", ip.IsUnspecified())
	}
}
