package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// IsMorse проверяет, является ли строка кодом Морзе
func IsMorse(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '\t' && r != '\n' {
			return false
		}
	}
	return true
}

// Convert автоматически определяет тип данных и конвертирует их
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", nil
	}

	if IsMorse(input) {
		result := morse.ToText(input)
		return result, nil
	} else {
		result := morse.ToMorse(input)
		return result, nil
	}
}
