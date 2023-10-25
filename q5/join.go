package main

func Join(strs []string, sep string) string {

	s := ""
	length := 0
	for i := range strs {
		length = i + 1
	}
	length--
	for i, ele := range strs {
		s += string(ele)
		if i < length {
			s += sep
		}
	}
	return s
}
