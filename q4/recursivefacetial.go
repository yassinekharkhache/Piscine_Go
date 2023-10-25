package main

func RecursiveFactorial(nb int) int {
	s := 1
	if nb < 0{
		s = -s
		nb = -nb
	}
	if(nb > 0){
		nb *= RecursiveFactorial(nb - 1);
	}else{
		return 1
	}
	return nb * s
}
