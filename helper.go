package main

import (
	"fmt"
	"os"
)

func exitWithErrorCode(errorCode int, message string) {
	fmt.Println(message)
	os.Exit(errorCode)
}
