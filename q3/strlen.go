package main

import "fmt"

func StrLen(s string) int {
	length := 0
	for i := range s {
		length = i + 1
	}
	fmt.Println(length)
	return length
}

func main() {
	StrLen("yassine")
}
