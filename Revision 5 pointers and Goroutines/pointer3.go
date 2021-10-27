package main

import "fmt"

type parent struct {
	name string
	age  int
}

func main() {
	//pointer to struct
	a := parent{
		name: "Hill",
		age:  49,
	}
	b := &a
	fmt.Println(b)
	fmt.Println(b.name)
	fmt.Println(b.age)
	//we can also modify the name
	b.name = "vill"
	fmt.Println(b)
}
