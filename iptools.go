package main

import (
	"math"
	"strconv"
	"strings"
)

func parseipv4Octets(ip string) [4]int {
	octetsAsStrings := strings.Split(ip, ".")
	if len(octetsAsStrings) != 4 {
		panic("invalid ip format")
	}

	octetsAsInts := [4]int{}
	for i := 0; i < len(octetsAsStrings); i++ {
		octetsAsInts[i], _ = strconv.Atoi(octetsAsStrings[i])
	}

	return octetsAsInts
}

func ipv4OctetsToDecimal(octetsAsInts [4]int) int {
	return octetsAsInts[0]*int(math.Pow(256, 3)) +
		octetsAsInts[1]*int(math.Pow(256, 2)) +
		octetsAsInts[2]*256 +
		octetsAsInts[3]
}
