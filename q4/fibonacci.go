package main

func Fibonacci(index int) int {
        if index < 0{
                return -1
        }
        if index == 1 || index == 0{
                return index
        }
        return Fibonacci(index - 1) + Fibonacci(index - 2)
}
