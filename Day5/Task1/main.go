package main

import "fmt"

func even(EvenChan chan int) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			EvenChan <- i
		}
	}
	close(EvenChan)
}

func odd(Oddchan chan int) {
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			Oddchan <- i
		}
	}
	close(Oddchan)
}

func main() {
	Evenchan := make(chan int)
	OddChan := make(chan int)

	go even(Evenchan)
	go odd(OddChan)

	fmt.Println("Even numbers : ")
	for n := range Evenchan {
		fmt.Println(n)
	}
	fmt.Println("Odd numbers : ")
	for n := range OddChan {
		fmt.Println(n)
	}
}
