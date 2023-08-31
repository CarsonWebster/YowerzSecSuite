package main

import (
	"flag"

	"github.com/CarsonWebster/YowerzSecSuite/pkgs/passwordchecker"
)

func main() {
	// fmt.Println("Hello, World!")
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

	if password != "" {
		passwordchecker.CheckPasswordStrength(password)
	}
}
