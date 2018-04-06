package main

import (
	"fmt"

)
// START OMIT
func main() {
  course:= "Go fundamentals "
	fmt.Println("\n Course is ", course)
  // pass address of variable
  changeCourse(&course)
  fmt.Println("\n Course is ", course)
}

// takes a pointer reference to the address passed
func changeCourse (course *string) string {
  *course = "a New Go Fundamentals"
  fmt.Println("\n changed course is ", *course)
  return *course
}
// END OMIT
