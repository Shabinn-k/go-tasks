package main

import "fmt"

func greet(name string,numb int) {
	fmt.Println("Welcome",name,"your number is",numb)
}

func add(a,b int)int{
	// fmt.Println(a+b)
	return a+b
}

func main() {
	func(){fmt.Println("anonymous function")}()
	greet("Shabin",10)
	add(10,90)
	sum:=add(5,5)
	fmt.Println(sum)
}
