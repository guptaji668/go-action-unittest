package main

import (
    "fmt"
)

func Add(a, b int) int {
    return a + b
}

func main() {
    

    result := Add(3,2)
    fmt.Printf("Result: %d\n", result)
}