package main

import (
	"flag"
	"fmt"
	"os"

	base64 "github.com/CarsonWebster/YowerzSecSuite/pkgs/encode"
	"github.com/CarsonWebster/YowerzSecSuite/pkgs/passwordchecker"
)

func main() {
	// fmt.Println("Hello, World!")
	// argsWithProg := os.Args
	// argsWithoutProg := os.Args[1:]
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)

	passwordCheckFlag := flag.String("password-check", "", "Check the strength of a password string")
	encodeBase64Flag := flag.String("encode-base64", "", "Encode a string in base64")
	decodeBase64Flag := flag.String("decode-base64", "", "Decode a string from base64")

	flag.Parse()

	// PASSWORD CHECKER
	if *passwordCheckFlag != "" {
		var passwordStrength = passwordchecker.CheckPasswordStrength(*passwordCheckFlag)
		fmt.Printf("The password %s has a strength of:\n", *passwordCheckFlag)
		fmt.Printf("	Length: %d\n", passwordStrength["length"])
		fmt.Printf("	Complexity: %d\n", passwordStrength["complexity"])
		fmt.Printf("	Entropy: %d\n", passwordStrength["entropy"])
		fmt.Printf("	Variance: %d\n", passwordStrength["variance"])
		fmt.Printf("	Overall: %d\n", passwordStrength["overall"])
		os.Exit(0)
	}

	// BASE64 ENCODER
	if *encodeBase64Flag != "" {
		fmt.Printf("%s\n", base64.Encode(*encodeBase64Flag))
		os.Exit(0)
	}
	if *decodeBase64Flag != "" {
		decoded, err := base64.Decode(*decodeBase64Flag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", decoded)
		os.Exit(0)
	}

}
