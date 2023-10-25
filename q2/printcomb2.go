package main

import "github.com/01-edu/z01"

func PrintComb2() {

	a := 0;
	b := 0;
	for( a <= 99 ){
		b = a + 1
		for b <= 99{
			z01.PrintRune( rune((a / 10) + 48)  )
			z01.PrintRune( rune((a % 10) + 48)  )
			z01.PrintRune(' ')
			z01.PrintRune( rune((b / 10) + 48)  )
			z01.PrintRune( rune((b % 10) + 48)  )
			if(a == 98){
				break;
			}
			z01.PrintRune(' ')
			z01.PrintRune(',')
			b++;
		}
		a++;
	}
	z01.PrintRune('\n')

}

func main(){
	PrintComb2();
}
