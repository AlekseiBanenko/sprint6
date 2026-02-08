package service

import (
	"errors"
	"strings"

	"github.com/AlekseiBanenko/sprint6/pkg/morse"
)

func DetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("пустая строка")
	}

	// Проверка, является ли строка морзе-кодом
	if isMorseCode(trimmed) {
		// Конвертируем из морзе в текст
		return morse.ToText(trimmed), nil
	} else {
		// Конвертируем из текста в морзе
		return morse.ToMorse(trimmed), nil
	}
}

// Вспомогательная функция для определения типа строки
func isMorseCode(s string) bool {
	// Морзе состоит из точек, тире, пробелов
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '/' {
			return false
		}
	}
	return true
}
