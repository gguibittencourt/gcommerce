package validator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	err := RegisterValidation()
	assert.Nil(t, err)

	type entity struct {
		Required string `validate:"required"`
		Number   string `validate:"numeric"`
		ValidCPF string `validate:"validCPF"`
	}

	tests := []struct {
		name   string
		entity interface{}
		err    error
	}{
		{
			name:   "given an valid entity, should not return error",
			entity: entity{Required: "value1", Number: "1200", ValidCPF: "470.508.590-60"},
			err:    nil,
		},
		{
			name:   "given an invalid entity, should return error",
			entity: entity{Required: "", Number: "abc", ValidCPF: ""},
			err: fmt.Errorf(
				"Key: 'entity.Required' Error:Field validation for 'Required' failed on the 'required' tag" +
					"\nKey: 'entity.Number' Error:Field validation for 'Number' failed on the 'numeric' tag" +
					"\nKey: 'entity.ValidCPF' Error:Field validation for 'ValidCPF' failed on the 'validCPF' tag",
			),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Validate(test.entity)

			if test.err != nil {
				require.EqualError(t, err, test.err.Error())
				return
			}

			require.NoError(t, err)
		})
	}
}

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

func TestRegisterValidationMap(t *testing.T) {
	validFunc := func(validator.FieldLevel) bool { return true }
	tests := []struct {
		name        string
		tag         string
		validation  func(validator.FieldLevel) bool
		expectedErr bool
	}{
		{
			name:        "given a valid entry, should not return error",
			tag:         "isValid",
			validation:  validFunc,
			expectedErr: false,
		},
		{
			name:        "given an invalid entry, should return error",
			tag:         "",
			validation:  validFunc,
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vMap := map[string]func(validator.FieldLevel) bool{
				test.tag: test.validation,
			}
			err := registerValidationMap(vMap)
			assert.Equal(t, test.expectedErr, err != nil)
		})
	}
}
