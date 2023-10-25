package main

import "github.com/01-edu/z01"

func main(){
	x := 'a'
	for x <= 'z'{
		z01.PrintRune(x);
		x += 1;
	}
	z01.PrintRune('\n');

}
