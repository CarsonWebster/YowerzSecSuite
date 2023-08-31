package main

import (
	"flag"
	"fmt"

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
		var passwordStrength = passwordchecker.CheckPasswordStrength(password)
		fmt.Printf("The password %s has a strength of:\n", password)
		fmt.Printf("	Length: %d\n", passwordStrength["length"])
		fmt.Printf("	Complexity: %d\n", passwordStrength["complexity"])
		fmt.Printf("	Entropy: %d\n", passwordStrength["entropy"])
		fmt.Printf("	Variance: %d\n", passwordStrength["variance"])
		fmt.Printf("	Overall: %d\n", passwordStrength["overall"])
	}
}
