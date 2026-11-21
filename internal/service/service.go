package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет тип входных данных:
//   - если это код Морзе -> конвертирует в текст;
//   - иначе (обычный текст) -> конвертирует в код Морзе.
func Convert(input string) (string, error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return "", errors.New("пустые входные данные")
	}

	if isMorse(s) {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}

// isMorse возвращает true, если строка состоит только из символов кода Морзе:
// точки, тире, пробелов и переводов строк.
func isMorse(s string) bool {
	for _, r := range s {
		switch r {
		case '.', '-', ' ', '\n', '\r', '\t':
			// допустимые символы
		default:
			return false
		}
	}
	return true
}
