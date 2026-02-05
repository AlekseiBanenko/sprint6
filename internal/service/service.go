package service

import (
	"errors"
	"strings"
)

var (
	morseToText = map[string]string{
		".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E",
		"..-.": "F", "--.": "G", "....": "H", "..": "I", ".---": "J",
		"-.-": "K", ".-..": "L", "--": "M", "-.": "N", "---": "O",
		".--.": "P", "--.-": "Q", ".-.": "R", "...": "S", "-": "T",
		"..-": "U", "...-": "V", ".--": "W", "-..-": "X", "-.--": "Y",
		"--..": "Z", ".----": "1", "..---": "2", "...--": "3", "....-": "4",
		".....": "5", "-....": "6", "--...": "7", "---..": "8", "----.": "9",
		"-----": "0",
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
	}
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	if strings.ContainsAny(trimmed, ".-") {
		result, err := morseToTextDecode(trimmed)
		if err == nil {
			return result, nil
		}
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
				result.WriteByte(text[0])
			} else {
				return "", errors.New("invalid morse code")
			}
		}
	}
	return result.String(), nil
}

func textToMorseEncode(input string) string {
	var result strings.Builder
	for i, r := range input {
		if i > 0 && r == ' ' {
			result.WriteString(" / ")
			continue
		}
		if morse, ok := textToMorse[r]; ok {
			if i > 0 {
				result.WriteByte(' ')
			}
			result.WriteString(morse)
		}
	}
	return result.String()
}
