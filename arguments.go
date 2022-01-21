package main

import (
	"os"
	"strings"
)

func hasArgument(shortArgument string, longArgument string) bool {
	for i := 1; i < len(os.Args); i++ {
		cArgument := strings.Trim(os.Args[i], " ")
		if cArgument == "--"+shortArgument || cArgument == "-"+longArgument {
			return true
		}
	}

	return false
}

func getArgumentValue(shortArgument string, longArgument string, defaultValue string) string {

	for i := 1; i < len(os.Args); i++ {
		cArgument := strings.Trim(os.Args[i], " ")

		strings.IndexRune(cArgument, '=')

		if cArgument == "--"+shortArgument || cArgument == "-"+longArgument {
			//	return true
		}
	}

	return defaultValue
}

func parseOptions() Options {
	options := Options{}

	// input
	options.ipv4Address = strings.Trim(os.Args[1], " ")
	options.ipv4Octets = ipv4StringToOctets(options.ipv4Address)

	// options
	options.outputType = getArgumentValue("output", "o", "")

	return options
}
