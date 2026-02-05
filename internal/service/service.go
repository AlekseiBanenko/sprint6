package service

import (
	"errors"
	"strings"
)

var (
	morseToText = map[string]string{
		".-": "А", "−...": "Б", "−..−.": "В", "−..": "Г", ".": "Д",
		"..−.": "Е", "--.": "Ё", "....": "Ж", "..": "З", ".−−−": "И",
		"−.−": "Й", "−.": "К", ".−..": "Л", "--": "М", "−.": "Н",
		"−−−": "О", ".−−.": "П", "--−.": "Р", ".−.": "С", "−": "Т",
		"..−": "У", "...−": "Ф", ".−−": "Х", "−..−": "Ц", "−.−−": "Ч",
		"−−..": "Ш", "−−.−": "Щ", "−.−−−": "Ъ", "−−−−": "Ы", "−.−.": "Ь",
		"−−−.": "Э", ".−..−": "Ю", "..−−.": "Я",
	}

	textToMorse = map[rune]string{
		'А': ".-", 'Б': "−...", 'В': "−..−.", 'Г': "−..", 'Д': ".",
		'Е': "..−.", 'Ё': "--.", 'Ж': "....", 'З': "..", 'И': ".−−−",
		'Й': "−.−", 'К': "−.", 'Л': ".−..", 'М': "--", 'Н': "−.",
		'О': "−−−", 'П': ".−−.", 'Р': "--−.", 'С': ".−.", 'Т': "−",
		'У': "..−", 'Ф': "...−", 'Х': ".−−", 'Ц': "−..−", 'Ч': "−.−−",
		'Ш': "−−..", 'Щ': "−−.−", 'Ъ': "−.−−−", 'Ы': "−−−−", 'Ь': "−.−.",
		'Э': "−−−.", 'Ю': ".−..−", 'Я': "..−−.", ' ': "/",
	}
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// ✅ ФИКС: проверяем НАЛИЧИЕ ТОЧКИ-ТИРЕ
	if strings.ContainsAny(trimmed, "−") || strings.ContainsAny(trimmed, ".−") {
		return morseToTextDecode(trimmed)
	}
	return textToMorseEncode(trimmed), nil
}

func morseToTextDecode(input string) (string, error) {
	words := strings.Split(input, " / ")
	var result strings.Builder

	for i, word := range words {
		if i > 0 {
			result.WriteByte(' ')
		}
		symbols := strings.Split(word, " ")
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

func textToMorseEncode(input string) string {
	var result strings.Builder
	prevSpace := false

	for i, r := range input {
		if r == ' ' {
			if !prevSpace {
				if result.Len() > 0 {
					result.WriteString(" / ")
				}
				prevSpace = true
			}
			continue
		}
		prevSpace = false

		if morse, ok := textToMorse[r]; ok {
			if result.Len() > 0 {
				result.WriteByte(' ')
			}
			result.WriteString(morse)
		}
	}
	return result.String()
}
