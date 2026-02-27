// palindrome
// package main
// import "fmt"
// func main() {
// 	str := "malayalamo"
// 	r := []rune(str)
// 	f := 0
// 	l := len(str) - 1
// 	p := true
// 	for f < l {
// 		if r[f] != r[l] {
// 			p=false
// 			break
// 		}
// 		f++
// 		l--
// 	}
// 	fmt.Println(p)
// }

// You are not good nigga just do print the hello world first then learn golanbg blud
// package main
// import "fmt"
// func main() {
// 	str:=" You are not good nigga just do print the hello world first then learn golanbg blud"
// 	v:=0
// 	c:=0
// 	for i:=0;i<len(str);i++{
// 		ch:=str[i]

// 		if ch=='A'||ch=='a'||ch=='E'||ch=='e'||ch=='I'||ch=='i'||ch=='O'||ch=='o'||ch=='U'||ch=='u'{
// 			v++
// 		}else if ch !=' '{
// 			c++
// 		}
// 	}
// 	fmt.Println("Vowels : ",v)
// 	fmt.Println("consonants :",c)
// }

// square of even and cube of odd using channel
// package main
// import "fmt"
// func even(evench chan int) {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			evench <- i
// 		}
// 	}
// 	close(evench)
// }
// func odd(oddch chan int){
// 	for i:=1;i<=10;i++{
// 		if i%2!=0{
// 			oddch<-i
// 		}
// 	}
// 	close(oddch)
// }
// func main() {
// 	evench := make(chan int)
// 	oddch := make(chan int)
// 	go even(evench)
// 	go odd(oddch)
// 	fmt.Println("Even square :")
// 	for n:=range evench{
// 		fmt.Println(n,n*n)
// 	}
// 	fmt.Println("Odd cube : ")
// 	for n:=range oddch{
// 		fmt.Println(n,n*n*n)
// 	}
// }

// sum of even and odd
// package main
// import "fmt"
// func main() {
// 	e:=0
// 	o:=0
// 	for i:=0;i<=10;i++{
// 		if i%2==0{
// 			e+=i
// 		}else{
// 			o+=i
// 		}
// 	}
// 	fmt.Println("even : ",e," odd : ",o)
// }

// prevent a race condition
// package main
// import ("fmt"
// "sync")
// var count int
// var mu sync.Mutex
// func trial(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	for i:=0;i<1000;i++{
// 		mu.Lock()
// 		count++
// 		mu.Unlock()
// 	}
// }
// func main() {
// 	var wg sync.WaitGroup
// 	for i:=1;i<=5;i++{
// 		wg.Add(1)
// 		go trial(&wg)
// 	}
// 	wg.Wait()
// 	fmt.Println(count)
// }

// reverse a string
// package main
// import "fmt"
// func main() {
// 		str:="Hello World"
// 		f:=0
// 		l:=len(str)-1
// 		r:=[]rune(str)
// 		for f<l{
// 			r[f],r[l]=r[l],r[f]
// 			f++
// 			l--
// 		}
// 		fmt.Println(string(r))
// }

// store a string's index and char in a map
// package main
// import "fmt"
// import "fmt"
// func main() {
// 	str:="Hello World"
// 	m:=make(map[int]string)
// 	for i:=0;i<len(str);i++{
// 		m[i]=string(str[i])
// 	}
// 	fmt.Println(m)
// }

// fibonacci
// package main
// import "fmt"
// func main() {
// 	a:=0
// 	b:=1
// 	for i:=0;i<10;i++{
// 		fmt.Println(a)
// 		next:=a+b
// 		a=b
// 		b=next
// 	}
// }

// error if its is not even
// package main
// import (
// 	"errors"
// 	"fmt"
// )
// func check(n int)error{
// 	if n%2 !=0{
// 		return errors.New("It is odd number !")
// 	}
// 	return nil
// }
// func main() {
// 	var num int
// 	fmt.Println("Enter a number")
// 	fmt.Scan(&num)

// 	err:=check(num)
// 	if err!=nil{
// 		fmt.Println(err)``
// 	}else{
// 		fmt.Println("It is Even number !")
// 	}
// }

//panic recover
// package main
// import "fmt"
// func risk(){
// 	defer func(){
// 		if r:=recover();r!=nil{
// 			fmt.Println("recovered from",r)
// 		}
// 	}()
// 	panic("Panic occured")
// }
// func main() {
// 	risk()
// }

// interface
// package main

// import "fmt"
// type Shape interface{
// 	Rectangle() int
// 	Circle() float64
// 	Square() int
// }
// type Area struct{
// 	Width int
// 	Height int
// 	Radius float64
// }
// func (a Area)Rectangle()int{
// 	return a.Height*a.Width
// }
// func (a Area)Circle()float64{
// 	return 2*3.149890*a.Radius
// }
// func (a Area)Square()int{
// 	return a.Width*a.Width
// }
// func Get(s Shape){
// 	fmt.Println(s.Rectangle())
// 	fmt.Println(s.Circle())
// 	fmt.Println(s.Square())
// }
// func main() {
// 	size:=Area{Width:10 ,Height:25, Radius:4.5 }
// 	Get(size)
// }

// largest number
// package main
// import "fmt"
// func main() {
// 	nums:=[]int{3,4,25,105,0,12,53,76,99}
// 	max:=0
// 	for i:=0;i<len(nums);i++{
// 		if nums[i]>max{
// 			max=nums[i]
// 		}
// 	}
// 	fmt.Println(max)
// }

// remove duplicates form an array
// package main
// import "fmt"
// func main() {
// 	arr:=[]int{1,2,2,3,1,43,5,23,23,6,10,9,9,3,87,52}
// 	m:=make(map[int]int)
// 	dup:=[]int{}
// 	for _,n:=range arr{
// 		m[n]++
// 		if m[n]==2{
// 			dup=append(dup,n)
// 		}
// 	}
// 	fmt.Println(dup)
// }

package main

func main() {
	
}