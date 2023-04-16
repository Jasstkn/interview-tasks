package validator

import (
	"regexp"
	"strconv"
	"strings"
)

func hasLeadingZeros(input string) bool {
	return input[0] == '0' && len(input) != 1
}

func IPValidator(input string) (bool, error) {

	inputOctets := strings.Split(input, ".")

	// invalid if input is empty string or has >= 4 octets
	if len(inputOctets) != 4 {
		return false, nil
	}

	for _, octet := range inputOctets {
		if hasLeadingZeros(octet) {
			return false, nil
		}

		octetInt, err := strconv.Atoi(octet)
		if err != nil {
			return false, err
		}

		if octetInt < 0 || octetInt > 255 {
			return false, nil
		}
	}

	return true, nil
}

func IPValidatorRegexp(input string) bool {
	pattern := "^(((1[0-9]{0,2})|(2[0-5]{0,2})|0)[.]){3}((1[0-9]{0,2})|(2[0-5]{0,2})|0)$"
	result, ok := regexp.MatchString(pattern, input)
	if ok != nil {
		panic("invalid regex: " + pattern)
	}
	return result
}
