package main

func IsUpper(s string) bool {
        
        for _,ele := range s {
                if !(ele >= 'A' && ele <= 'Z'){
                        return false
                }
        }
        return true
}
