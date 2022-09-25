package validator

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	firstVerifierDigitPosition  = 10
	secondVerifierDigitPosition = 11
)

func IsValidCPF(value string) bool {
	if value == "" {
		return false
	}
	digits := cleanNonDigits(value)
	if !isValidLength(digits) || equalDigits(digits) {
		return false
	}
	cpfToValidate := extractDigit(digits)
	firstDigit := calculateVerifierDigit(cpfToValidate, firstVerifierDigitPosition)
	cpfToValidate += firstDigit
	secondDigit := calculateVerifierDigit(cpfToValidate, secondVerifierDigitPosition)
	return digits == cpfToValidate+secondDigit
}

func equalDigits(digits string) bool {
	firstDigit := digits[:1]
	return strings.ReplaceAll(digits, firstDigit, "") == ""
}

func isValidLength(cpf string) bool {
	return len(cpf) == 11
}

func extractDigit(cpf string) string {
	return cpf[:9]
}

func calculateVerifierDigit(cpf string, position int) string {
	var sum int
	for _, r := range cpf {
		sum += toInt(r) * position
		position--
	}
	sum %= 11
	if sum < 2 {
		return "0"
	}
	return strconv.Itoa(11 - sum)
}

func cleanNonDigits(s string) string {
	digitsRegex := regexp.MustCompile(`\D`)
	return digitsRegex.ReplaceAllString(s, "")
}

func toInt(r rune) int {
	return int(r - '0')
}
