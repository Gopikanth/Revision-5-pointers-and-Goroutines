package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan string)
	b := make(chan string)
	go channel(a)
	go channel2(b)
	select {
	case option1 := <-a:
		fmt.Println(option1)
	case option2 := <-b:
		fmt.Println(option2)
	}

}
func channel(other chan string) {
	time.Sleep(2 * time.Microsecond)
	other <- "I LOVE TRICHY "
}
func channel2(other2 chan string) {
	time.Sleep(5 * time.Microsecond)
	other2 <- "Chennai"

}
