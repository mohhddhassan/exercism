package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	replacer := strings.NewReplacer(
		"-", "",
		".", "",
		" ", "",
		"+1", "",
		"(", "",
		")", "",
	)

	result := replacer.Replace(phoneNumber)
    
    if len(result) == 11 && strings.HasPrefix(result, "1") {
    result = result[1:]
    }

	if len(result) != 10 {
		return "", fmt.Errorf("phone number should be 10 digits")
	}

	for _, r := range result {
		if !unicode.IsDigit(r) || r > '9' {
			return "", fmt.Errorf("phone number should contain only digits")
		}
	}

	if result[0] < '2' || result[0] > '9' {
		return "", fmt.Errorf("phone number should start with a digit from 2 to 9")
	}

	if result[3] < '2' || result[3] > '9' {
		return "", fmt.Errorf("exchange code should start with a digit from 2 to 9")
	}

	return result, nil
}

func AreaCode(phoneNumber string) (string, error) {
	ac, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return ac[:3], nil
}

func Format(phoneNumber string) (string, error) {
	fr, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", fr[:3], fr[3:6], fr[6:]), nil
}