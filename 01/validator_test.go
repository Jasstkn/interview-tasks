package validator

import "testing"

func TestIPValidator(t *testing.T) {
	testCases := []struct {
		ip    string
		valid bool
	}{
		{"", false},
		{"192.168.0.A", false},
		{"192.168.0.1", true},
		{"256.0.0.1", false},
		{"001.0.0.1", false},
		{"100.001.0.1", false},
		{"100.1.0.1.123", false},
		{"100.1.0", false},
	}

	for _, testCase := range testCases {
		isValid, _ := IPValidator(testCase.ip)
		if isValid != testCase.valid {
			t.Errorf("Result for %s was incorrect, got: %t, want: %t.", testCase.ip, isValid, testCase.valid)
		}
	}
}
