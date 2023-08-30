package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)

	passwordFlag := flag.String("P", "", "Check password strength")
	aliasPasswordFlag := flag.String("password-check", "", "Alias for -P (Check password strength)")
	flag.Parse()

	// Combine the values from both flags
	password := *passwordFlag
	if password == "" {
		password = *aliasPasswordFlag
	}

	if *passwordFlag != "" {
		fmt.Println("Password is", password)
	} else {
		fmt.Println("Usage: ./YowerzSecSuite [-P <password>] [--password-check <password>]")
		os.Exit(1)
	}
}
