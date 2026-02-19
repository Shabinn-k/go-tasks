package main
import "fmt"

func main() {

	array:=[]int{1,2,3,4}
	var arr [3]int
	fmt.Println(arr)
	fmt.Println(len(array))

	//slice
	s := []int{10, 20, 30}
	s=append(s, 40)
	fmt.Println(s)

	//append variadic
	a:=[]string{"a","b","c"}
	b:=[]string{"d","e","f"}
	a=append(a,b...)
	fmt.Println(a)
}