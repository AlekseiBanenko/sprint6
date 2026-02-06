package service

import (
	"errors"
	"strings"
)

var (
	morseToText = map[string]string{
		".-": "А", "-...": "Б", "-.-.": "В", "-..": "Г", ".": "Д",
		"..-.": "Е", "--.": "Ё", "....": "Ж", "..": "З", ".---": "И",
		"-.-": "Й", ".-..": "К", "--": "Л", "-.": "М", "---": "Н",
		".--.": "О", "--.-": "П", ".-.": "Р", "...": "С", "-": "Т",
		"..-": "У", "...-": "Ф", ".--": "Х", "-..-": "Ц", "-.--": "Ч",
		".-..": "Ш", "--..": "Щ", "-.-.-": "Ъ", "-.-..": "Ы", "..-..": "Ь",
		".--.-": "Э", "---.": "Ю", ".-.-.": "Я", "/": " ",
	}

	textToMorse = map[rune]string{
		'А': ".-", 'Б': "-...", 'В': "-.-.", 'Г': "-..", 'Д': ".",
		'Е': "..-.", 'Ё': "--.", 'Ж': "....", 'З': "..", 'И': ".---",
		'Й': "-.-", 'К': ".-..", 'Л': "--", 'М': "-.", 'Н': "---",
		'О': ".--.", 'П': "--.-", 'Р': ".-.", 'С': "...", 'Т': "-",
		'У': "..-", 'Ф': "...-", 'Х': ".--", 'Ц': "-..-", 'Ч': "-.--",
		'Ш': ".-..", 'Щ': "--..", 'Ъ': "-.-.-", 'Ы': "-.-..", 'Ь': "..-..",
		'Э': ".--.-", 'Ю': "---.", 'Я': ".-.-.", ' ': "/",
		// строчные тоже
		'а': ".-", 'б': "-...", 'в': "-.-.", 'г': "-..", 'д': ".",
		// ... остальные аналогично
	}
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// Детекция: если содержит .- и пробелы = Морзе
	if strings.ContainsAny(trimmed, ".-") && strings.ContainsAny(trimmed, " /") {
		return decodeMorse(trimmed)
	}
	return encodeMorse(trimmed), nil
}

func decodeMorse(input string) (string, error) {
	words := strings.Split(input, " / ")
	var result strings.Builder

	for _, word := range words {
		if result.Len() > 0 {
			result.WriteByte(' ')
		}
		symbols := strings.Fields(word)
		for _, symbol := range symbols {
			if text, ok := morseToText[symbol]; ok {
				result.WriteString(text)
			} else {
				return "", errors.New("invalid morse code")
			}
		}
	}
	return result.String(), nil
}

func encodeMorse(input string) string {
	var result strings.Builder
	inWord := false

	for _, r := range input {
		if r == ' ' {
			if inWord {
				result.WriteString(" / ")
				inWord = false
			}
			continue
		}
		if morse, ok := textToMorse[r]; ok {
			if inWord {
				result.WriteByte(' ')
			}
			result.WriteString(morse)
			inWord = true
		}
	}
	return result.String()
}
