// Fibanocci series with a loop
package main

import "fmt"

func main() {
	var num1 int
	a := 0
	b := 1
	fmt.Println("Enter a number :")
	fmt.Scan(&num1)
	for i := 0; i <= num1; i++ {
		fmt.Println(a)
		next := a + b
		a = b
		b = next
	}
}
