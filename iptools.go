package main

import (
	"math"
	"strconv"
	"strings"
)

// converts ipv4 address string to numeric octets
func ipv4StringToOctets(ip string) [4]byte {
	octets := strings.Split(ip, ".")
	if len(octets) != 4 {
		panic("invalid ip format")
	}

	octetsAsByte := [4]byte{}
	for i := 0; i < len(octets); i++ {
		v, _ := strconv.Atoi(octets[i])
		octetsAsByte[i] = byte(v)
	}

	return octetsAsByte
}

// converts ipv4 numeric octets to decimal representation
func ip4vOctetsToDecimal(octets [4]byte) int {
	v1 := int(octets[0]) * int(math.Pow(256, 3))
	v2 := int(octets[1]) * int(math.Pow(256, 2))
	v3 := int(octets[2]) * 256
	v4 := int(octets[3])

	return v1 + v2 + v3 + v4
}
