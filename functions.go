package main
import (
	"fmt"
	"strings"
)

func main() {
	module := "functions basics"
	author := "vikash singh"
	mod, auth, error := converter(module, author)
  if(error == nil) {
    fmt.Println(mod, auth)
  }
}

func converter(module, author string) (s1, s2 string, err error) {
	module = strings.Title(module)
	author = strings.ToUpper(author)
  // nil indicates success
	return module, author, nil
}
