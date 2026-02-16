// Simple calculator using switch.
package main

import "fmt"

func main() {
	var num1, num2 int
	var prop string

	fmt.Println("Enter two  Number : ")
	fmt.Scan(&num1, &num2)
	fmt.Println("Which method ?")
	fmt.Scan(&prop)

	switch prop {
	case "+":
		fmt.Println("Your answer : ", num1+num2)
	case "-":
		fmt.Println("Your answer is : ", num1-num2)
	case "*":
		fmt.Println("Your answer is : ", num1*num2)
	case "/":
		fmt.Println("Your answer is : ", num1/num2)
	default:
		fmt.Println("Dont be oversmart!")
	}
}
