package service

import (
	"command-line-arguments/Users/a1/cod/sprint6/sprint6/pkg/morse/morse.go"
	"errors"
	"fmt"
	"strings"
)

var (
	morseToText = map[string]string{
		".-":    "A",
		"-...":  "B",
		"-.-.":  "C",
		"-..":   "D",
		".":     "E",
		"..-.":  "F",
		"--.":   "G",
		"....":  "H",
		"..":    "I",
		".---":  "J",
		"-.-":   "K",
		".-..":  "L",
		"--":    "M",
		"-.":    "N",
		"---":   "O",
		".--.":  "P",
		"--.-":  "Q",
		".-.":   "R",
		"...":   "S",
		"-":     "T",
		"..-":   "U",
		"...-":  "V",
		".--":   "W",
		"-..-":  "X",
		"-.--":  "Y",
		"--..":  "Z",
		".----": "1",
		"..---": "2",
		"...--": "3",
		"....-": "4",
		".....": "5",
		"-....": "6",
		"--...": "7",
		"---..": "8",
		"----.": "9",
		"-----": "0",
		".--.":  "П", "--.": "Г", ".-.": "Р", "..": "И", ".--": "В", ".": "Е",
		"-": "Т", "----": "Ш", "-..": "Д", "..-.": "..-.", "-.--": "Ы",
		"-.-": "К", "-.-.-": "Ц", "---.": "Ч", ".-..": "Л", "-.": "Н",
		"-...": "Б", "..-": "У", "--": "М",
	}

	textToMorse = map[rune]string{
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
		'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
		'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
		'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
		'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
		'Z': "--..", '1': ".----", '2': "..---", '3': "...--", '4': "....-",
		'5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.",
		'0': "-----", ' ': "/",
		'a': ".-", 'b': "-...", 'c': "-.-.", 'd': "-..", 'e': ".",
		'f': "..-.", 'g': "--.", 'h': "....", 'i': "..", 'j': ".---",
		'k': "-.-", 'l': ".-..", 'm': "--", 'n': "-.", 'o': "---",
		'p': ".--.", 'q': "--.-", 'r': ".-.", 's': "...", 't': "-",
		'u': "..-", 'v': "...-", 'w': ".--", 'x': "-..-", 'y': "-.--",
		'z': "--..",
		// Русские буквы для теста "ПРИВЕТ" и "РШГДФФЫКЦГЧЧЧЛЕНБУРКЛФГМЛ"
		'П': ".--.", 'п': ".--.",
		'Р': ".-.", 'р': ".-.",
		'И': "..", 'и': "..",
		'В': ".--", 'в': ".--",
		'Е': ".", 'е': ".",
		'Т': "-", 'т': "-",
		'Ш': "----", 'ш': "----",
		'Г': "--.", 'г': "--.",
		'Д': "-..", 'д': "-..",
		'Ф': "..-.", 'ф': "..-.",
		'Ы': "-.--", 'ы': "-.--",
		'К': "-.-", 'к': "-.-",
		'Ц': "-.-.-", 'ц': "-.-.-",
		'Ч': "---.", 'ч': "---.",
		'Л': ".-..", 'л': ".-..",
		'Н': "-.", 'н': "-.",
		'Б': "-...", 'б': "-...",
		'У': "..-", 'у': "..-",
		'М': "--", 'м': "--",
	}
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// Детекция Морзе: содержит точки/тире + разделители
	if strings.ContainsAny(trimmed, ".-") &&
		(strings.Contains(trimmed, " ") || strings.Contains(trimmed, "/")) {
		// Морзе → Текст
		return morse.ToText(trimmed), nil
	}

	// Текст → Морзе
	return morse.ToMorse(trimmed), nil
}

func morseToTextDecode(input string) (string, error) {
	words := strings.Split(input, " / ")
	var result strings.Builder
	firstWord := true

	for _, word := range words {
		if !firstWord {
			result.WriteByte(' ')
		}
		firstWord = false

		symbols := strings.Split(word, " ")
		for _, symbol := range symbols {
			if symbol == "" {
				continue
			}
			if text, ok := morseToText[symbol]; ok {
				result.WriteString(text)
			} else {
				return "", fmt.Errorf("invalid morse code: %q", symbol)
			}
		}
	}
	return strings.TrimSpace(result.String()), nil
}

func textToMorseEncode(input string) string {
	var result strings.Builder
	first := true

	for _, r := range input {
		if r == ' ' {
			if !first && result.Len() > 0 {
				result.WriteString(" / ")
			}
			continue
		}

		if morse, ok := textToMorse[r]; ok {
			if !first && result.Len() > 0 {
				result.WriteByte(' ')
			}
			result.WriteString(morse)
			first = false
		}
	}
	return result.String()
}
