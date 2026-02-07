package service

import (
	"errors"
	"fmt"
	"strings"
)

var (
	morseToText = map[string]string{
		".-":     "А",
		"-...":   "Б",
		"-.-.":   "В",
		"-..":    "Г",
		".":      "Д",
		"..-.":   "Е",
		"--.":    "Ё",
		"....":   "Ж",
		"..":     "З",
		".---":   "И",
		"-.-":    "Й",
		".-..":   "К", // ← УНИКАЛЬНЫЙ!
		"--":     "Л",
		"-.":     "М",
		"---":    "Н",
		".--.":   "О",
		"--.-":   "П",
		".-.":    "Р",
		"...":    "С",
		"-":      "Т",
		"..-":    "У",
		"...-":   "Ф",
		".--":    "Х",
		"-..-":   "Ц",
		"-.--":   "Ч",
		"----":   "Ш",
		"--..":   "Щ",
		"-.-.-":  "Ъ",
		"-.-..":  "Ы",
		"..-..":  "Ь",
		".--.-":  "Э",
		"---.":   "Ю",
		".-.-.-": "Я",
		"/":      " ",
	}

	textToMorse = map[rune]string{
		'А': ".-", 'а': ".-",
		'Б': "-...", 'б': "-...",
		'В': "-.-.", 'в': "-.-.",
		'Г': "-..", 'г': "-..",
		'Д': ".", 'д': ".",
		'Е': "..-.", 'е': "..-.",
		'Ё': "--.", 'ё': "--.",
		'Ж': "....", 'ж': "....",
		'З': "..", 'з': "..",
		'И': ".---", 'и': ".---",
		'Й': "-.-", 'й': "-.-",
		'К': ".-..", 'к': ".-..",
		'Л': "--", 'л': "--",
		'М': "-.", 'м': "-.",
		'Н': "---", 'н': "---",
		'О': ".--.", 'о': ".--.",
		'П': "--.-", 'п': "--.-",
		'Р': ".-.", 'р': ".-.",
		'С': "...", 'с': "...",
		'Т': "-", 'т': "-",
		'У': "..-", 'у': "..-",
		'Ф': "...-", 'ф': "...-",
		'Х': ".--", 'х': ".--",
		'Ц': "-..-", 'ц': "-..-",
		'Ч': "-.--", 'ч': "-.--",
		'Ш': "----", 'ш': "----",
		'Щ': "--..", 'щ': "--..",
		'Ъ': "-.-.-", 'ъ': "-.-.-",
		'Ы': "-.-..", 'ы': "-.-..",
		'Ь': "..-..", 'ь': "..-..",
		'Э': ".--.-", 'э': ".--.-",
		'Ю': "---.", 'ю': "---.",
		'Я': ".-.-.-", 'я': ".-.-.-",
		' ': "/",
	}
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// Проверка, содержит ли вход только допустимые символы для Морзе
	morseChars := ".- /"
	isMorse := true
	for _, r := range trimmed {
		if !strings.ContainsRune(morseChars, r) && r != ' ' {
			isMorse = false
			break
		}
	}

	if isMorse {
		return decodeMorse(trimmed)
	}
	return encodeMorse(trimmed), nil
}

func decodeMorse(input string) (string, error) {
	words := strings.Split(input, " / ")
	var result strings.Builder

	for i, word := range words {
		if i > 0 {
			result.WriteByte(' ')
		}
		for _, symbol := range strings.Fields(word) {
			if text, ok := morseToText[symbol]; ok {
				result.WriteString(text)
			} else {
				return "", fmt.Errorf("invalid morse code: %q", symbol)
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
