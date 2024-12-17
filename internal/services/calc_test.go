package services

import (
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		expected    float64
		shouldError bool
	}{
		{"Addition", "3 + 5", 8, false},
		{"Subtraction", "10 - 2", 8, false},
		{"Multiplication", "2 * 3", 6, false},
		{"Division", "8 / 2", 4, false},
		{"Division by zero", "5 / 0", 0, true},
		{"Parentheses", "(3 + 5) * 2", 16, false},
		{"Invalid character", "3 + a", 0, true},
		{"Empty expression", "", 0, true},
		{"Unbalanced parentheses", "(3 + 5", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Calc(test.expression)
			if test.shouldError {
				if err == nil {
					t.Errorf("Expected error, got nil for expression: %s", test.expression)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for expression: %s, error: %v", test.expression, err)
				}
				if result != test.expected {
					t.Errorf("For expression: %s, expected %f, got %f", test.expression, test.expected, result)
				}
			}
		})
	}
}
