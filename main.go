package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Options struct {
	ipAddress       string
	outputAsDecimal bool
}

type IpInformations struct {
	decimalValue int
}

func main() {
	argsCount := len(os.Args)
	if argsCount < 2 {
		exitWithErrorCode(1, "Syntax: cidr <ip-address> [options]")
	} else {
		options := parseOptions()
		processIp(options)
	}
}

func exitWithErrorCode(errorCode int, message string) {
	fmt.Println(message)
	os.Exit(errorCode)
}

func containsArgument(searchValue string, alias string) bool {
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--"+searchValue || os.Args[i] == "-"+alias {
			return true
		}
	}

	return false
}

func parseOptions() Options {
	options := Options{}
	options.ipAddress = strings.Trim(os.Args[1], " ")
	options.outputAsDecimal = containsArgument("decimal", "d")

	return options
}

func processIp(options Options) {
	octetsAsInts := parseipv4Octets(options.ipAddress)

	informations := IpInformations{}
	informations.decimalValue = ipv4OctetsToDecimal(octetsAsInts)

	writeOutput(options, informations)
}

func writeOutput(options Options, informations IpInformations) {
	if options.outputAsDecimal {
		fmt.Println(informations.decimalValue)
	} else {
		fmt.Println("Decimal", informations.decimalValue)
	}
}

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
