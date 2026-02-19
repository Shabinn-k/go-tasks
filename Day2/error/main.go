package main

import (
	"errors"
	"fmt"
)

func checkNumb(num1 int) error {
	if num1%2 < 10 {
		return errors.New("greater than 10")
	}
	return nil
}
func main() {
	errr:=fmt.Errorf("such a good error")
	fmt.Println(errr)
	err := checkNumb(100)
	if err != nil {
		fmt.Println("Error", err)
	}
}
