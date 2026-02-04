package service

import (
	"errors"
	"strings"
)

var morseCode = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.",
	'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.",
	'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..", '0': "-----", '1': ".----", '2': "..---",
	'3': "...--", '4': "....-", '5': ".....", '6': "-....", '7': "--...",
	'8': "---..", '9': "----.", ' ': "/",
}

var englishCode = map[string]rune{
	".-": 'A', "-...": 'B', "-.-.": 'C', "-..": 'D', ".": 'E',
	"..-.": 'F', "--.": 'G', "....": 'H', "..": 'I', ".---": 'J',
	"-.-": 'K', ".-..": 'L', "--": 'M', "-.": 'N', "---": 'O',
	".--.": 'P', "--.-": 'Q', ".-.": 'R', "...": 'S', "-": 'T',
	"..-": 'U', "...-": 'V', ".--": 'W', "-..-": 'X', "-.--": 'Y',
	"--..": 'Z', "-----": '0', ".----": '1', "..---": '2', "...--": '3',
	"....-": '4', ".....": '5', "-....": '6', "--...": '7', "---..": '8',
	"----.": '9', "/": ' ',
}

func AutoConvert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("input is empty")
	}

	words := strings.Split(input, " ")
	if len(words) == 0 || (strings.HasPrefix(words[0], ".") || strings.HasPrefix(words[0], "-")) {
		return morseToEnglish(input)
	}
	return textToMorse(input)
}

func textToMorse(text string) (string, error) {
	var result []string
	for _, char := range strings.ToUpper(text) {
		if code, ok := morseCode[char]; ok {
			result = append(result, code)
		}
	}
	return strings.Join(result, " "), nil
}

func morseToEnglish(morse string) (string, error) {
	words := strings.Split(morse, "/")
	var result strings.Builder
	for i, word := range words {
		letters := strings.Split(word, " ")
		for _, letter := range letters {
			if char, ok := englishCode[letter]; ok {
				result.WriteRune(char)
			}
		}
		if i < len(words)-1 {
			result.WriteRune(' ')
		}
	}
	return result.String(), nil
}
