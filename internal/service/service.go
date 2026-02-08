package service

import (
	"errors"
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// DetectAndConvert определяет, является ли входной текст морзе-кодом,
// и конвертирует его либо в текст, либо в морзе.
unc DetectAndConvert(input string) (string, error) {
    trimmed := strings.TrimSpace(input)

    log.Printf("DetectAndConvert: входные данные='%s'", trimmed)

    if trimmed == "" {
        return "", errors.New("пустая строка")
    }

    // Проверка, является ли строка JSON
    if strings.HasPrefix(trimmed, "{") || strings.HasPrefix(trimmed, "[") {
        log.Println("Обнаружен JSON, возвращаем как есть.")
        return trimmed, nil
    }

    // Проверка, является ли строка морзе-кодом
    if isMorseCode(trimmed) {
        log.Println("Обнаружен морзе-код, конвертация в текст.")
        return morse.ToText(trimmed), nil
    } else {
        log.Println("Обнаружен обычный текст, конвертация в морзе.")
        return morse.ToMorse(trimmed), nil
    }
}


// Вспомогательная функция для определения, является ли строка морзе-кодом
func isMorseCode(s string) bool {
	hasMorseChars := false
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '/' {
			return false
		}
		if r == '.' || r == '-' {
			hasMorseChars = true
		}
	}
	return hasMorseChars
}
