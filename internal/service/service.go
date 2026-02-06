package service

import (
	"errors"
	"strings"

	"github.com/AlekseiBanenko/sprint6/pkg/morse"
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
