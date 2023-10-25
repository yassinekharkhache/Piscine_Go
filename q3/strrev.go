package main

func StrRev(s string) string {
	length := 0
	for i := range s {
		length = i + 1
	}
	length--
	r := ""
	for i := length; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}
