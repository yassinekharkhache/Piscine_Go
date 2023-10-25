package main

func NRune(s string, n int) rune {
	length := 0
	for i := range s {
		length = i + 1
	}
	if n > length || n < 0 {
		return 0
	}
	return rune(s[n-1])
}
