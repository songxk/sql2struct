package main

import "strings"

func ucLower(str string) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return str
	}
	if len(str) == 1 {
		return strings.ToLower(str)
	}
	firstLetter := str[:1]
	return strings.ToLower(firstLetter) + str[1:]
}

func ucUpper(str string) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return str
	}
	if len(str) == 1 {
		return strings.ToUpper(str)
	}
	firstLetter := str[:1]
	return strings.ToUpper(firstLetter) + str[1:]
}
