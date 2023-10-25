package main

func IsAlpha(s string) bool {
        
        for _,ele := range s {
                if !(ele >= 'a' && ele <= 'z' || ele >= 'A' && ele <= 'Z' || ele >= '0' && ele <= '9'){
                        return false
                }
        }
        return true
}
