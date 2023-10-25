package main

func LastRune(s string) rune {
	length := 0
	for i := range s {
		length = i + 1
	}
	length--
	return rune(s[length])
}
