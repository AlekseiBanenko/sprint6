package service

import (
	"errors"
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
	}
)

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// Если содержит последовательности точек-тире разделённые пробелами = Morse
	if strings.ContainsAny(trimmed, ". -") {
		return morseToTextDecode(trimmed)
	}

	// Текст → Morse
	return textToMorseEncode(trimmed), nil
}

func morseToTextDecode(input string) (string, error) {
	words := strings.Split(input, " / ")
	var result strings.Builder

	for _, word := range words { // Fixed: _ instead of i
		symbols := strings.Split(word, " ")
		for _, symbol := range symbols {
			if text, ok := morseToText[symbol]; ok {
				result.WriteString(text)
			} else {
				return "", errors.New("invalid morse code")
			}
		}
		if len(words) > 1 && result.Len() > 0 {
			result.WriteByte(' ')
		}
	}
	return strings.TrimSpace(result.String()), nil
}

func textToMorseEncode(input string) string {
	var result strings.Builder

	for i, r := range input {
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
