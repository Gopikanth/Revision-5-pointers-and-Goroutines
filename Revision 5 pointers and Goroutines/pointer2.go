package main

import "fmt"

func main() {
//pasing pointer to function
	var a int = 100
	fmt.Println("THE VALUE NOW IS a = ", a)
	p := &a
	ptr(p)
	fmt.Println("THE VALUE after becomes as = ", a)
	c := 300
	fmt.Println("value of c = ", c)
	pt(&c)
}
func ptr(b *int) {
	*b = 200
}

//passing address to function
func pt(d *int) {
	*d = 400
}
