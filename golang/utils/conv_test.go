package utils

import (
	"testing"
)

func TestConvertBinaryStringToDecimal(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"1100", 12},
		{"100", 4},
		{"0", 0},
		{"1", 1},
		{"11111111", 255},
		{"111111111", 511},
	}
	for _, tc := range cases {
		actual := ConvertBinaryStringToDecimal(tc.input)
		if actual != tc.expected {
			t.Fatalf("expected: %d got: %d for input: %s",
				tc.expected, actual, tc.input)
		}
	}
}
