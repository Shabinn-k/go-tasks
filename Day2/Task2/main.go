package main

import (
	"fmt"
	"golang/Day2/Task2/arithmetic"
)

func main() {
	fmt.Println("Add : ",arithmetic.Add(10,6))
	fmt.Println("Substract : ",arithmetic.Substract(100,1))
	fmt.Println("Multiply : ",arithmetic.Multiply(10,10))
result, err := arithmetic.Divide(110, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide:", result)
	}	
}