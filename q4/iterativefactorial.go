package main

import "fmt"

func IterativeFactorial(nb int) int {
	r := 1
	s := 1
	if nb < 0{
		nb = -nb
		s = -s

	}
	for nb > 0{
		r *= nb;
		nb--;
	}
	return r*s
}
func main() {
	arg := -4
	fmt.Println(IterativeFactorial(arg))
}
