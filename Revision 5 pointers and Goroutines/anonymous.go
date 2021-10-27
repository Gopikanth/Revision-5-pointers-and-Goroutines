package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Desktops")

	go func() {
		fmt.Println("Laptops")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Android phones")
}
