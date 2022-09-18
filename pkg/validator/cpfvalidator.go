package validator

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

const (
	firstVerifierDigitPosition  = 10
	secondVerifierDigitPosition = 11
)

var (
	cpfRegexp = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)
)

func IsValidCPF(cpf string) bool {
	if cpf == "" {
		return false
	}
	if !cpfRegexp.MatchString(cpf) {
		return false
	}
	digits := cleanNonDigits(cpf)
	cpfToValidate := digits[:9]
	firstDigit := calculateVerifierDigit(cpfToValidate, firstVerifierDigitPosition)
	cpfToValidate += firstDigit
	secondDigit := calculateVerifierDigit(cpfToValidate, secondVerifierDigitPosition)
	return digits == cpfToValidate+secondDigit
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
	buf := bytes.NewBufferString("")
	for _, r := range s {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func toInt(r rune) int {
	return int(r - '0')
}
