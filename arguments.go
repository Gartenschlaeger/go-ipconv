package main

import (
	"os"
	"strings"
)

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
