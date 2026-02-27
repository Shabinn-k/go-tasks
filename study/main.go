//Reverse a String
// package main
// import "fmt"
// func main() {
// 	str:="hello world"
// 	r:=[]rune(str)
// 	f:=0
// 	l:=len(str)-1
// 	for f<l{
// 		r[f],r[l]=r[l],r[f]
// 		f++
// 		l--
// 	}
// 	fmt.Println(string(r))
// }

//Count Vowels
// package main
// import "fmt"
// func main() {
// 	str:="hello world"
// 	v:=0
// 	c:=0
// 	for i:=0;i<len(str);i++{
// 		ch:=str[i]
// 		if ch=='a'||ch=='A'||ch=='e'||ch=='E'||ch=='i'||ch=='I'||ch=='o'||ch=='O'||ch=='u'||ch=='U'{
// 			v++
// 		}else{
// 			c++
// 		}
// 	}
// 	fmt.Println("vowels -",v)
// 	fmt.Println("Consonants -",c)
// }

//find larget number
// package main
// import "fmt"
// func main() {
// 	arr:= []int{10, 20, 4, 45, 99}
// 	max:=0
// 	min:=0
// 	for i:=0;i<len(arr);i++{
// 		if arr[i]>max{
// 		max=arr[i]
// 		}else{
// 			min=arr[i]
// 		}
// 	}
// 	fmt.Println(max,"-",min)
// }

//Check Palindrome
// package main
// import "fmt"
// func main() {
// 	p:=true
// 	str:="malayalam"
// 	r:=[]rune(str)
// 	f:=0
// 	l:=len(str)-1
// 	for f<l{
// 		if r[f]!=r[l]{
// 			p=false
// 			break
// 		}
// 		f++
// 		 l--
// 	}
// 	fmt.Println(p)
// }

//Remove Duplicates from Slice
// package main
// import "fmt"
// func main() {
// 	arr:=[]int{1,2,3,12,3,1,1,3,53,52,2,5,53,100}
// 	m:=make(map[int]int)
// 	plane:=[]int{}
// 	for _,val:=range arr{
// 		m[val]++
// 		if m[val]==1{
// 			plane = append(plane, val)
// 	}
// }
// 	dupe:=[]int{}
// 	og:=[]int{}
// 	for k,c:=range m{
// 		if c>1{
// 			dupe=append(dupe, k)
// 		}else{
// 			og=append(og, k)
// 		}
// 	}
// 	fmt.Println("Array ",arr)
// 	fmt.Println("removed spam",plane)
// 	fmt.Println("duplicates",dupe)
// 	fmt.Println("not duplicates",og)
// }

// repeating words
// package main
// import (
// 	"fmt"
// 	"strings"
// )
// func main() {
// 	str := `then I asked him with my eyes to ask again yes and then he
// 	        asked me would I yes to say yes my mountain flower and first
// 	        I put my arms around him yes and drew him down to me so he 
// 	        could feel my breasts all perfume yes and his heart was
// 	        going like mad and yes I said yes I will Yes`
// 	words:=strings.Fields(str)
// 	m:=make(map[string]int)
// 	for _,val:=range words{
// 		m[val]++
// 	}
// 	fmt.Println(m)
// }

//
package main
import "fmt"
type Student struct{
	Name string
	Score []int
}
func(s Student)Average()float64{
	if len(s.Score)==0{
		return 0
	}
	sum:=0
	for n:=range s.Score{
		sum+=n
	}
	return float64(sum)/float64(len(s.Score))
}
func main() {
	st:=Student{
		Name: "shabin",
		Score:[]int{10,20,30,40,50,60,80},
	}
	fmt.Println("Student",st.Name)
	fmt.Println(st.Average())
}