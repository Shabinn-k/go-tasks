package main

import "fmt"

func math(a, b int) (int, int, int, int) {
	sum := a + b
	diff := a - b
	prod := a * b
	quot := a / b
	return sum, diff, prod, quot
}

func main() {
	s, d, p, q := math(100, 10)
	fmt.Println("sum = ", s)
	fmt.Println("Difference = ", d)
	fmt.Println("product = ", p)
	fmt.Println("quotient = ", q)
}
