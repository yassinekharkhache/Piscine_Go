package main

func ToLower(s string) string {
         str := []byte(s)
        for i,ele := range str{
                if( ele > 'A' && ele < 'Z'){
                        str[i] = str[i] + 32
                }
        }
        return string(str)
}
