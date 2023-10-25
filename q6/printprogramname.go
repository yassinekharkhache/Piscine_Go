package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printstr(str string){
	for _,ele := range str{
		z01.PrintRune(ele)
	}
}


func main(){
	x := os.Args[0]
	printstr(x);
}



