package main

import "fmt"

type Student struct {
	Name  string
	Grade int
	Age   int
}

func main() {
	studentInfo := map[string]Student{
		"shabin": {Age: 18, Grade: 123},
		"john":   {Age: 11, Grade: 100},
	}
	fmt.Println(studentInfo["shabin"].Age)
	fmt.Println(studentInfo["john"].Grade)
}
