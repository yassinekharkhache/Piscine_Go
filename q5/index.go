package main

func Index(s string, toFind string) int {

	length := 0
	for i := range s {
		length = i + 1
	}
	leng := 0
	for i := range toFind {
		leng = i + 1
	}
	leng--
	for i := 0; i < length; i++ {
		for j := 0; s[i+j] == toFind[j]; j++ {
			if j == (leng) {
				return (i)
			}
		}
	}
	return (-1)
}
