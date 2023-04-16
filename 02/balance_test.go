package balance

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	ip    string
	valid bool
}{
	{"0.0.0.1", true}
}

func TestIPValidator(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("IPValidator/%s", testCase.ip), func(t *testing.T) {
			isValid, _ := IPValidator(testCase.ip)
			if isValid != testCase.valid {
				t.Errorf("Result for %s was incorrect, got: %t, want: %t.", testCase.ip, isValid, testCase.valid)
			}
		})
		t.Run(fmt.Sprintf("IPValidatorRegexp/%s", testCase.ip), func(t *testing.T) {
			isValid := IPValidatorRegexp(testCase.ip)
			if isValid != testCase.valid {
				t.Errorf("Result for %s was incorrect, got: %t, want: %t.", testCase.ip, isValid, testCase.valid)
			}
		})
	}
}

// benchmarks
func BenchmarkIPValidator(b *testing.B) {
	for _, testCase := range testCases {
		IPValidator(testCase.ip)
	}
}

func BenchmarkIPValidatorRegexp(b *testing.B) {
	for _, testCase := range testCases {
		IPValidatorRegexp(testCase.ip)
	}
}
