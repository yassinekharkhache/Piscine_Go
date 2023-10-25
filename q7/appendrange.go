package main

func AppendRange(min, max int) []int {
        var s = max - min;
         if s < 0{
                return nil
        }
        var arr []int;
        for i:=0;i<s;i++{
            arr = append(arr,min);
            min++;
        }
        return arr
}
