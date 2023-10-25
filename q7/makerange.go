package main

func MakeRange(min, max int) []int {
        s := max - min
        if s < 0{
                return nil
        }
        arr := make([]int,s)
        for i := 0; i < s;i++ {
                arr[i] = min
                min++;
        }
        return arr
}
