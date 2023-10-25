package main

func Any(f func(string) bool, a []string) bool {
	for _, ele := range a {

		if f(string(ele)) == true {
			return true
		}
	}
	return false

}
