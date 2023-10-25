package main

func ToUpper(s string) string {
        str := []byte(s)
        for i,ele := range str{
                if( ele > 'a' && ele < 'z'){
                        str[i] = str[i] - 32
                }
        }
        return string(str)
}
