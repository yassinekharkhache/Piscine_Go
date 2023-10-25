package main

func Capitalize(s string) string {
	s = ToLower(s)
	length := 0
	for i := range s {
		length = i + 1
	}
	length--

	str := []byte(s)
	for i, ele := range str {
		if !((ele >= 'a' && ele <= 'z') || (ele >= '0' && ele <= '9') || (ele >= 'A' && ele <= 'Z')) && i < length {
			str[i+1] = str[i+1] - 32

		}
	}
	if str[0] >= 'a' && str[0] <= 'z' {
		str[0] -= 32
	}
	return string(str)

}

func ToLower(s string) string {
	str := []byte(s)
	for i, ele := range str {
		if ele >= 'A' && ele <= 'Z' {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}
