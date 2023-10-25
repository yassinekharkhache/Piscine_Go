package main

import "github.com/01-edu/z01"

func main(){
	x := 'z'

	for x >= 'a'{
		z01.PrintRune(x);
		x--;
	}
	z01.PrintRune('\n');
}
