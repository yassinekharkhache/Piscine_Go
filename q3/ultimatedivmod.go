package main 

func UltimateDivMod(a *int, b *int) {
	x := *a / *b;
	*b = *a % *b;
	*a = x;
}
