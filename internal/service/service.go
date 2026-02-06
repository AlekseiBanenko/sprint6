package service

import (
	"errors"
	"strings"

	"github.com/AlekseiBanenko/sprint6/pkg/morse" // ← Путь к вашему модулю
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// Детекция Морзе: содержит . и - + разделители
	if strings.Contains(trimmed, ".") &&
		strings.Contains(trimmed, "-") &&
		(strings.Contains(trimmed, " ") || strings.Contains(trimmed, "/")) {
		// Морзе → Текст
		return morse.ToText(trimmed), nil
	}

	// Текст → Морзе
	return morse.ToMorse(trimmed), nil
}
