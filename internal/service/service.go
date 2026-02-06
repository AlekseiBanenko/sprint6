package service

import (
	"errors"
	"pkg/morse" // Локальный пакет из шаблона
	"strings"
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// Если содержит . - и разделители = Морзе
	if strings.ContainsAny(trimmed, ".-") {
		return morse.ToText(trimmed), nil
	}

	// Текст → Морзе
	return morse.ToMorse(trimmed), nil
}
