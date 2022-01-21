package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/apparentlymart/go-cidr/cidr"
)

func convertCidr(cidr string) string {
	r := regexp_ipv4cidr.FindStringSubmatch(cidr)

	octets := []string{r[1], r[2], r[3], r[4], r[5]}
	for i := 0; i < len(octets); i++ {
		if octets[i] == "" {
			octets[i] = ".0"
		}
	}

	return strings.Join(octets, "")
}

func processIpv4Cidr(ipv4Cidr string) {
	convertedCidr := convertCidr(ipv4Cidr)

	_, ipNet, err := net.ParseCIDR(convertedCidr)
	if err == nil {
		firstIp, lastIp := cidr.AddressRange(ipNet)
		addressCount := cidr.AddressCount(ipNet)

		fmt.Printf("cidr               %v\n", convertedCidr)

		fmt.Printf("netmask            %v.%v.%v.%v\n",
			ipNet.Mask[0],
			ipNet.Mask[1],
			ipNet.Mask[2],
			ipNet.Mask[3])

		fmt.Printf("address count      %v\n", addressCount)

		fmt.Printf("network address    %v\n", firstIp)

		fmt.Printf("broadcast address  %v\n", lastIp)

		if addressCount > 2 {
			fmt.Printf("first host         %v\n", cidr.Inc(firstIp))
			fmt.Printf("last host          %v\n", cidr.Dec(lastIp))
		} else {
			fmt.Printf("first host         %v\n", firstIp)
			fmt.Printf("last host          %v\n", lastIp)
		}
	} else {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
