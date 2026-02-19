package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	students := []Student{
		{Name: "shabin", Age: 18,},
		{Name: "john", Age: 20,},
	}
	for _, value := range students {
		fmt.Println(value.Name,"-",value.Age)
	}
}