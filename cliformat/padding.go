package cliformat

import "strings"

func LeftPad(text string, length int, ch string) string {
	if length-len(text) < 0 {
		return text
	}
	return strings.Repeat(ch, length-len(text)) + text
}

func RightPad(text string, length int, ch string) string {
	if length-len(text) < 0 {
		return text
	}
	return text + strings.Repeat(ch, length-len(text))
}
