package main

func Map(f func(int) bool, a []int) []bool {

	length := 0
	for i := range a {
		length = i + 1
	}
	var x = make([]bool, length)

	for i, ele := range a {
		x[i] = f(ele)
	}
	return x
}
