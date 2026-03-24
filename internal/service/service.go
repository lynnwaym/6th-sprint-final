package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseOrText(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", fmt.Errorf("Empty string")
	}

	if isMorseCode(input) {
		result := morse.ToText(input)
		if result == "" {

			return "", fmt.Errorf("convert to text failed")
		}
		return result, nil
	}
	result := morse.ToMorse(input)
	if result == "" {

		return "", fmt.Errorf("convert to morse failed")
	}
	return result, nil
}

func isMorseCode(s string) bool {

	for _, ch := range s {
		if ch != '.' && ch != '-' && ch != ' ' {
			return false
		}
	}
	return true
}
