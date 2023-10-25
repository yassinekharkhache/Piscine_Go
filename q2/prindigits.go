package main

import "github.com/01-edu/z01"

func main(){
        x := '0'

        for x <= '9'{
                z01.PrintRune(x);
                x++;
        }
        z01.PrintRune('\n');
}

