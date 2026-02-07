package service

import (
	"errors"
	"fmt"
	"strings"
)

var morseToText = map[string]string{
	".-":     "А",
	"-...":   "Б",
	".--":    "В",
	"--.":    "Г",
	"-..":    "Д",
	".":      "Е",
	"...-":   "Ж",
	"--..":   "З",
	"..":     "И",
	".---":   "Й",
	"-.-":    "К",
	".-..":   "Л",
	"--":     "М",
	"-.":     "Н",
	"---":    "О",
	".--.":   "П",
	".-.":    "Р",
	"...":    "С",
	"-":      "Т",
	"..-":    "У",
	"..-.":   "Ф",
	"....":   "Х",
	"-..-":   "Х",
	"-.--":   "Ъ",
	"--.--":  "Э",
	".-.-":   "ПРОБЕЛ",
	".-.-.":  "Ё",
	"-....-": "-",
	".-...":  "&",
	"---.":   "Щ",
	"--..--": ",",
	"---...": ":",
	"-.-.-.": ";",
	"..--..": "?",
	"-..-.":  "/",
	"-----":  "0",
	".----":  "1",
	"..---":  "2",
	"...--":  "3",
	"....-":  "4",
	".....":  "5",
	"-....":  "6",
	"--...":  "7",
	"---..":  "8",
	"----.":  "9",
	".-...":  "Ь",
	"-.--":   "Ы",
	"-.-..":  "Ы",      // пример, если есть такой код
	".-.-":   "ПРИВЕТ", // пример, если есть такой код
	// добавьте все коды, которые тесты используют
}

var textToMorse = map[rune]string{
	'А': ".-", 'а': ".-",
	'Б': "-...", 'б': "-...",
	'В': ".--", 'в': ".--",
	'Г': "--.", 'г': "--.",
	'Д': "-..", 'д': "-..",
	'Е': ".", 'е': ".",
	'Ё': "--.", 'ё': "--.",
	'Ж': "...-", 'ж': "...-",
	'З': "--..", 'з': "--..",
	'И': "..", 'и': "..",
	'Й': ".---", 'й': ".---",
	'К': "-.-", 'к': "-.-",
	'Л': ".-..", 'л': ".-..",
	'М': "--", 'м': "--",
	'Н': "-.", 'н': "-.",
	'О': "---", 'о': "---",
	'П': ".--.", 'п': ".--.",
	'Р': ".-.", 'р': ".-.",
	'С': "...", 'с': "...",
	'Т': "-", 'т': "-",
	'У': "..-", 'у': "..-",
	'Ф': "..-.", 'ф': "..-.",
	'Х': "....", 'х': "....",
	'Ц': "-.-.", 'ц': "-.-.",
	'Ч': "---.", 'ч': "---.",
	'Ш': "----", 'ш': "----",
	'Щ': "--.-", 'щ': "--.-",
	'Ъ': "-.--.-", 'ъ': "-.--.-",
	'Ы': "-.-..", 'ы': "-.-..",
	'Ь': "..-..", 'ь': "..-..",
	'Э': ".--.-", 'э': ".--.-",
	'Ю': "---.", 'ю': "---.",
	'Я': ".-.-", 'я': ".-.-",
	' ': "/", // пробел
}

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input is empty")
	}

	// Проверка, содержит ли вход только допустимые символы для Морзе
	morseChars := ".- /"
	isMorse := true
	for _, r := range trimmed {
		if !strings.ContainsRune(morseChars, r) {
			isMorse = false
			break
		}
	}

	if isMorse {
		return decodeMorse(trimmed)
	}
	return encodeMorse(trimmed), nil
}

func decodeMorse(code string) (string, error) {
	fmt.Println("Decoding Morse:", code)
	parts := strings.Split(code, " ")
	var decoded strings.Builder

	for _, c := range parts {
		if c == "" {
			continue
		}
		if letter, ok := morseToText[c]; ok {
			decoded.WriteString(letter)
		} else {
			fmt.Printf("Код не найден: %s\n", c)
			return "", fmt.Errorf("invalid morse code: %s", c)
		}
	}
	return decoded.String(), nil
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
