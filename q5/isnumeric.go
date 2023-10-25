package main

func IsNumeric(s string) bool {
        
        for _,ele := range s {
                if !(ele >= '0' && ele <= '9'){
                        return false
                }
        }
        return true
}
