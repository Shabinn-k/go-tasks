package main

import "fmt"

func main() {
	name := "Shabin"
	fmt.Println(name)
	age := 18
	fmt.Println(age)
	active := true
	fmt.Println(active)
	emoji := '😋'
	fmt.Println(string(emoji))
	var a float32 =1.1234567890
    b :=1.1234567890
    fmt.Println(a)
    fmt.Println(b)
    fmt.Printf("hyy my name is %s and im %d years old \n i %tly love to do works, i like %c pizza",name,age,active,emoji)
    size:=[]int{1,2,3,4,5}
    fmt.Println(len(size))
}
