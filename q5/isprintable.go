package main

func IsPrintable(s string) bool {

        for _,ele := range s{
                
                if !(ele >= 32 && ele <= 126){
                        return false
                }
        }
        return true
}
