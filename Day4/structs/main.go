package main

import "fmt"

type User struct {
	name string
	age  int
}

type Greet struct {
	name    string
	age     int
	salary  float64
	sencere bool
	emo     rune
}

func (u *User) birthday() {
    u.age++
}

func (g Greet)Welcome() {
	fmt.Printf("In the age of %d , %s is a very good hardworking man with a salary range of %flpa . he is %tly sencere to his work %c \n",
		g.age, g.name, g.salary, g.sencere, g.emo)
}

func main() {
	 
	user1 := User{
		name: "shabin",
		age:  18,
	}
	user1.name = "john"

	fmt.Println(user1.name)
	fmt.Println(user1.age)

	user2 := User{}
	fmt.Println(user2)
	fmt.Println(user1)

	Job := Greet{
		name:    "shabin",
		age:     18,
		salary:  20.91231381937912319999,
		sencere: true,
		emo:     '🌹',
	}
	Job.Welcome()
 
}