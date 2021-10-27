package main

import "fmt"

func main() {
	first := 123
	second := 456
	third := &first
	fourth := &second
	//comparing the pointers with == operator
	five := third == fourth
	fmt.Println(five)
	//we can also use != operator
	six :=&first != &second
	fmt.Println(six)

}