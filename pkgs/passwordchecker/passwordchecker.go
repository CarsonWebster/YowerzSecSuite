package passwordchecker

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

// PasswordStrength represents the strength evaluation of a password.
type PasswordStrength map[string]int

// CheckPasswordStrength checks the strength of a password
func CheckPasswordStrength(password string) PasswordStrength {
	// Number of password grading tests
	numTests := 0
	// Initialize strength map
	strength := make(PasswordStrength)

	// Grade based on length
	lengthGrade := len(password)
	if lengthGrade > 10 {
		lengthGrade = 10
	}
	strength["length"] = lengthGrade
	numTests++

	// Grade based on complexity (using a simple pattern here, adapt as needed)
	complexityGrade := calculateComplexityGrade(password)
	strength["complexity"] = complexityGrade
	numTests++

	// Grade based on entropy
	entropyGrade := calculateEntropyGrade(password)
	strength["entropy"] = entropyGrade
	numTests++

	// Grade based on variance
	varianceGrade := calculateVarianceGrade(password)
	strength["variance"] = varianceGrade
	numTests++

	// Grade based on sequences
	// sequencesGrade := calculateSequencesGrade(password)
	// strength["sequences"] = sequencesGrade
	// numTests++

	// Calculate overall grade
	overallGrade := 0
	for _, value := range strength {
		overallGrade += value
	}
	overallGrade = overallGrade / numTests
	strength["overall"] = overallGrade

	return strength
}

// calculateComplexityGrade calculates the complexity grade based on the password's complexity.
func calculateComplexityGrade(password string) int {
	// Regular expressions for different types of characters
	onlyNumbers, _ := regexp.MatchString("^[0-9]+$", password)
	onlyLowerCase, _ := regexp.MatchString("^[a-z]+$", password)
	onlyUpperCase, _ := regexp.MatchString("^[A-Z]+$", password)
	upperAndLowerCase, _ := regexp.MatchString("^[a-zA-Z]+$", password)
	upperLowerNumber, _ := regexp.MatchString("^[a-zA-Z0-9]+$", password)
	upperLowerNumberSymbol, _ := regexp.MatchString("^[a-zA-Z0-9!@#$%^&*()]+$", password)

	switch {
	case onlyNumbers:
		return 0
	case onlyLowerCase:
		return 2
	case onlyUpperCase:
		return 2
	case upperAndLowerCase:
		return 4
	case upperLowerNumber:
		return 6
	case upperLowerNumberSymbol:
		return 10
	default:
		return 0
	}
}

// calculateEntropyGrade calculates the entropy grade based on password's unpredictability.
func calculateEntropyGrade(password string) int {
	// Calculate entropy based on the number of possible characters and password length
	alphabetSize := 95 // Printable ASCII characters (excluding space)
	possibleCombinations := math.Pow(float64(alphabetSize), float64(len(password)))
	maxEntropy := math.Log2(possibleCombinations)
	normalizedEntropy := (maxEntropy / 10) // Normalize to 0-10 range, capped at 10

	entropyGrade := int(math.Round(math.Min(normalizedEntropy, 10.0)))
	return entropyGrade
}

func calculateVarianceGrade(password string) int {
	types := []func(rune) bool{
		unicode.IsLower,
		unicode.IsUpper,
		unicode.IsDigit,
		func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) },
	}
	counts := make([]bool, len(types))
	for _, r := range password {
		for i, charType := range types {
			if charType(r) {
				counts[i] = true
				break
			}
		}
	}
	varianceCount := 0
	for _, present := range counts {
		if present {
			varianceCount++
		}
	}

	// Normalize varianceCount to a score between 0 and 10
	varianceScore := (float64(varianceCount) / float64(len(types))) * 10
	return int(math.Round(varianceScore))
}

// calculateSequencesGrade calculates the sequences grade based on avoiding common sequences.
func calculateSequencesGrade(password string) int {
	sequences := []string{
		"0123456789", "9876543210",
		"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba",
		"abcdefghijklmnopqrstuvwxyz" + "abcdefghijklmnopqrstuvwxyz",
	}

	sequencesFound := 0
	for _, seq := range sequences {
		if strings.Contains(password, seq) {
			sequencesFound++
		}
	}

	return 10 - sequencesFound
}
