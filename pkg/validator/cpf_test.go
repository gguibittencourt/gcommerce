package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected bool
	}{
		{
			name:     "given a CPF valid, should return true",
			cpf:      "470.508.590-60",
			expected: true,
		},
		{
			name:     "given a CPF valid only numbers, should return true",
			cpf:      "47050859060",
			expected: true,
		},
		{
			name:     "given an empty CPF, should return false",
			cpf:      "",
			expected: false,
		},
		{
			name:     "given a CPF with invalid format, should return false",
			cpf:      "123",
			expected: false,
		},
		{
			name:     "given an invalid CPF with the same number, should return false",
			cpf:      "111.111.111-11",
			expected: false,
		},
		{
			name:     "given a CPF with invalid verifier digit, should return false",
			cpf:      "471.508.590-60",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsValidCPF(test.cpf)
			assert.Equal(t, test.expected, got)
		})
	}
}
