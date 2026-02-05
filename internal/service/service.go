package service

import (
	"errors"
	"strings"
)

var (
	morseToText = map[string]string{
		".-":    "А",
		"−...":  "Б",
		"−..−.": "В",
		"−..":   "Г",
		".":     "Д",
		"..−.":  "Е",
		"--.":   "Ё",
		"....":  "Ж",
		"..":    "З",
		".−−−":  "И",
		"−.−":   "Й",
		"−.-":   "К",
		".−..":  "Л",
		"--":    "М",
		"−.":    "Н",
		"−−−":   "О",
		".−−.":  "П",
		"--−.":  "Р",
		".−.":   "С",
		"−":     "Т",
		"..−":   "У",
		"...−":  "Ф",
		".−−":   "Х",
		"−..−":  "Ц",
		"−.−−":  "Ч",
		"−−..":  "Ш",
		"−−.−":  "Щ",
		"−.−−−": "Ъ",
		"−−−−":  "Ы",
		"−.−.":  "Ь",
		"−−−.":  "Э",
		".−..−": "Ю",
		"..−−.": "Я",
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

	if strings.ContainsAny(trimmed, "−") {
		return morseToTextDecode(trimmed)
	}
	return textToMorseEncode(trimmed), nil
}

func morseToTextDecode(input string) (string, error) {
	words := strings.Split(input, " / ")
	var result strings.Builder

	for _, word := range words {
		symbols := strings.Split(word, " ")
		for _, symbol := range symbols {
			if text, ok := morseToText[symbol]; ok {
				result.WriteString(text)
			} else {
				return "", errors.New("invalid morse code")
			}
		}
		result.WriteByte(' ')
	}
	return strings.TrimSpace(result.String()), nil
}

func textToMorseEncode(input string) string {
	var result strings.Builder

	for _, r := range input {
		if r == ' ' {
			if result.Len() > 0 {
				result.WriteString(" / ")
			}
			continue
		}

		if morse, ok := textToMorse[r]; ok {
			if result.Len() > 0 {
				result.WriteByte(' ')
			}
			result.WriteString(morse)
		}
	}
	return result.String()
}
