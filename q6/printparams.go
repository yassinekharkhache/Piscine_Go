package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printstr(str string) {
	for _, ele := range str {
		z01.PrintRune(ele)
	}
	z01.PrintRune('\n')
}

func main() {

	length := 0
	for i := range os.Args {
		length = i + 1
	}

	for i := 1; i < length; i++ {
		printstr(os.Args[i])
	}
}
