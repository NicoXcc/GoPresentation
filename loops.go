package main

import (
	"fmt"
	"time"
)
// For <expression>
//
// <expression>
// 1. Can be blank (infinite loop)
// 2. Can be a proper boolean expression
// 3. Can be a range
// START OMIT
func main() {
	// being idiomatic and initialized timer inside for loop
	for timer := 3; timer > 0; timer-- {
		fmt.Println("time left :", timer)
		time.Sleep(1 * time.Second)
	}
}
// END OMIT
