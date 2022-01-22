package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/c-robinson/iplib"
)

func processIpv4Address() {
	ip := net.ParseIP(*flag_input)
	if ip != nil {
		switch strings.ToLower(strings.TrimSpace(*flag_filter)) {
		default:
			fmt.Printf("address         %v\n", *flag_input)
			fmt.Printf("arpa            %v\n", iplib.IP4ToARPA(ip))
			fmt.Printf("is private      %v\n", ip.IsPrivate())
			fmt.Printf("is loopback     %v\n", ip.IsLoopback())
			fmt.Printf("is multicast    %v\n", ip.IsMulticast())
			fmt.Printf("is unspecified  %v\n", ip.IsUnspecified())
			fmt.Printf("dec             %v\n", iplib.IP4ToUint32(ip))
			fmt.Printf("bin             %v\n", iplib.IPToBinaryString(ip))
			fmt.Printf("hex             %v\n", iplib.IPToHexString(ip))
		case "arpa":
			fmt.Println(iplib.IP4ToARPA(ip))
		case "private":
			fmt.Println(ip.IsPrivate())
		case "loopback":
			fmt.Println(ip.IsLoopback())
		case "multicast", "broadcast":
			fmt.Println(ip.IsMulticast())
		case "unspecified":
			fmt.Println(ip.IsUnspecified())
		case "dec":
			fmt.Println(iplib.IP4ToUint32(ip))
		case "bin":
			fmt.Println(iplib.IPToBinaryString(ip))
		case "hex":
			fmt.Println(iplib.IPToHexString(ip))
		}
	} else {
		fmt.Println("Failed to parse input")
		os.Exit(1)
	}
}
