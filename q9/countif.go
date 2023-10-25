package main

func CountIf(f func(string) bool, tab []string) int {
	x := 0
	for _, ele := range tab {
		if f(ele) == true {
			x++
		}
	}
	return x
}
