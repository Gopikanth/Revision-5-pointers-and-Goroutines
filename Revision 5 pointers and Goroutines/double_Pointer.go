package main

import "fmt"

func main() {
	var a int = 45
	var b *int = &a
	var c **int = &b
	fmt.Println("The value of a = ", a)
	fmt.Println("The address of a", &a)
	fmt.Println("The value of b = ", b)
	fmt.Println("The address of b = ", &b)
	fmt.Println("The value of c = ", c)
	fmt.Println("The address of c = ", &c)
	fmt.Println("The value at address c is = ", *c)
	fmt.Println("The value at address c is = ", **c)
	//we can also alter the value
	**c = 56
	fmt.Println("The value of a now becomes as",**c)
}
