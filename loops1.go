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
	// declaring a slices
	course := []string{"math", "english", "hindi"}
	for _, i := range course {
		fmt.Println(i)
	}
}
// END OMIT
