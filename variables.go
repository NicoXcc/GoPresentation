package main
import (
	"fmt"
	"reflect"
)

var (
	name, course string
	module       float64
)
var (
	_name, _course, _module = "Vikash", "Go fundamentals", 3.2
)

func main() {
	fmt.Println("'name' is ", name, "and is of type ", reflect.TypeOf(name))
	fmt.Println("'module' is ", module, " and of type ", reflect.TypeOf(module))
	fmt.Println("'_name' is ", _name, "and is of type ", reflect.TypeOf(_name))
	fmt.Println("'_module' is ", _module, " and of type ", reflect.TypeOf(_module))
}
