package main

import "os"

func Atoi(s string) int {

	r := 0
	for _, ele := range s {
		if ele > '0' || ele < '9' {
			r *= 10
			r += int(ele) - 48

		}
	}
	return r
}

func main() {
	x := Atoi(os.Args[1])
	z := Atoi(os.Args[3])
	if os.Args[2] == "-" {
		println(x - z)
	} else if os.Args[2] == "*" {
		println(x * z)
	} else if os.Args[2] == "+" {
		println(x + z)
	} else if os.Args[2] == "/" {
		println(x / z)
	} else if os.Args[2] == "%" {
		println(x % z)
	}
}
