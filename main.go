package main

import (
    "fmt"
)

func Add(a, b int) int {
    return a + b
}

func main() {
    var num1, num2 int

    fmt.Print("Enter the first number: ")
    _, err1 := fmt.Scan(&num1)
    if err1 != nil {
        fmt.Println("Error reading the first number:", err1)
        return
    }

    fmt.Print("Enter the second number: ")
    _, err2 := fmt.Scan(&num2)
    if err2 != nil {
        fmt.Println("Error reading the second number:", err2)
        return
    }

    result := Add(num1, num2)
    fmt.Printf("Result: %d\n", result)
}