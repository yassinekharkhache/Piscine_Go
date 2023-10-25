package main

func BasicJoin(elems []string) string {
        s := ""
        for _,ele := range elems{
                
                s += string(ele)
        }
        return s
}
