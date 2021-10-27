package main

import (
	"fmt"
)

func main() {
	//variables storing hexadecimal values
	x := 0xFA
	fmt.Printf("Type is %T\n", x)
	fmt.Printf("value of hexadecimal is %X\n", x)
	fmt.Printf("value of decimal is %v\n", x)
	//Declaration and initialization of pointers
	var y int = 50
	p := &y
	fmt.Println("The value of y =", y)
	fmt.Println("The address of y =", &y)
	fmt.Println("The value of p =", p)
	//nil value of pointer
	var z *int
	fmt.Println(z) //the default value of pointer is nil
	//type inference in pointer variable
	var a = "gk" //we can use short hand also a:="gk"
	var q = &a   //taking pointer without type of variable
	fmt.Println("The value of z =", a)
	fmt.Println("The address of z=", &a)
	fmt.Println("The value of q =", q)
	//dereferencing the pointer
	var c int = 123
	b := &c
	fmt.Println(b)
	fmt.Println(*b) //in this we are dereferencing so that original value is displayed

	//we can also change the value of the pointer
	*b = 124
	fmt.Println(*b)

}
