package main
import (
	"fmt"
)
// For <expression>
//
// <expression>
// 1. Can be blank (infinite loop)
// 2. Can be a proper boolean expression
// 3. Can be a range
// START OMIT
func main() {
  x:= 3
	for x > 0 {
		fmt.Println("conditional looping")
		x--
	}

	for {
		fmt.Println("loop forever")
		break
	}
}
// END OMIT
