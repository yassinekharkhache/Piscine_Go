package main

func RecursivePower(nb int, power int) int {
        if power < 0 {
                return 0
        }
        if power > 0{
                nb *= RecursivePower(nb,power - 1)
        }else{
                return 1
        }
        return nb
}
