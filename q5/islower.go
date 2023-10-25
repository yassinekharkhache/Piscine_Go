package main

func IsLower(s string) bool {
        
        for _,ele := range s {
                if !(ele >= 'a' && ele <= 'z'){
                        return false
                }
        }
        return true
}
