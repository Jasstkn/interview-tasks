package validator

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	ip    string
	valid bool
}{
	{"0.0.0.1", true},
	{"192.168.0.1", true},
	{"255.255.255.255", true},
	{"", false},
	{"192.168.0.A", false},
	{"256.0.0.1", false},
	{"001.0.0.1", false},
	{"100.001.0.1", false},
	{"100.1.0.1.123", false},
	{"100.1.0", false},
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
