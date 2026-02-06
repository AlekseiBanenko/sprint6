package service

import (
	"errors"
	"pkg/morse" // ← ЛОКАЛЬНЫЙ ПАКЕТ!
	"strings"
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// Морзе содержит . - и разделители
	if strings.Contains(trimmed, ".") &&
		strings.Contains(trimmed, "-") &&
		(strings.Contains(trimmed, " ") || strings.Contains(trimmed, "/")) {
		return morse.ToText(trimmed), nil
	}

	return morse.ToMorse(trimmed), nil
}
