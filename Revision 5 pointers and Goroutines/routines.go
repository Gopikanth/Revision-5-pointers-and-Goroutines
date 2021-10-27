/*package main

import "fmt"

func main() {
a:= "How are you"
b := "i am fine"
go display(a) //here goroutines are not executed
display(b)
}
func display(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}

OUTPUT IS :
PS E:\Go programs\Revision 4> go run routines.go
i am fine
i am fine
i am fine
i am fine
i am fine	
*/
//now we can display the routine using sleep method()

package main

import (
	"fmt"
	"time"
)

func main() {
	a := "How are you"
	b := "i am fine"
	go display(a) //here goroutines are not executed
	display(b)
}
func display(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1*time.Second)
		fmt.Println(str)
	}
}
