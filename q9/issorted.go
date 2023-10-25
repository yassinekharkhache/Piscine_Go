package main

func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	length := 0
	for i := range a {
		length = i + 1
	}
	for i := 0; i < length-1; i++ {
		if f(a[i], a[i+1]) != -1 {
			return false
		}

	}
	return true
}
