package validator

import (
	"strconv"
	"strings"
)

func hasLeadingZeros(input string) bool {
	return input[0] == '0' && len(input) != 1
}

func IPValidator(input string) (bool, error) {

	inputOctets := strings.Split(input, ".")

	// invalid if input is empty string or has >= 4 octets
	if len(input) == 0 || len(inputOctets) != 4 {
		return false, nil
	}

	decision := true

	for _, octet := range inputOctets {
		if hasLeadingZeros(octet) {
			return false, nil
		}

		octetInt, err := strconv.Atoi(octet)
		if err != nil {
			return false, err
		}

		if octetInt < 0 || octetInt > 255 {
			decision = false
		}
	}

	return decision, nil
}
