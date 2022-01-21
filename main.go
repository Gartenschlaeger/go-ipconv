package main

import (
	"fmt"
	"os"
)

func main() {
	argsCount := len(os.Args)
	if argsCount < 2 {
		writeHelpOutput()
	} else {
		options := parseOptions()
		processIp(options)
	}
}

func processIp(options Options) {
	octetsAsInts := parseipv4Octets(options.ipAddress)

	informations := IpInformations{}
	informations.decimalValue = ipv4OctetsToDecimal(octetsAsInts)

	writeOptionsOutput(options, informations)
}

func writeHelpOutput() {
	exitWithErrorCode(1, "Syntax: cidr <ip-address> [options]")
}

func writeOptionsOutput(options Options, informations IpInformations) {
	if options.outputAsDecimal {
		fmt.Println(informations.decimalValue)
	} else {
		fmt.Println("Decimal", informations.decimalValue)
	}
}
