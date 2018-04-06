package main
import (
	"fmt"

)
// START OMIT
func main() {
  course:= "Go fundamentals "
	fmt.Println("\n Course is ", course)
  changeCourse(course)
}

func changeCourse (course string) string {
  course = "a New " + course
  fmt.Println("\n changed course is ", course)
  return course
}
// END OMIT
