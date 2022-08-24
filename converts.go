package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "sergio correa"
	course := "pepito"
	fmt.Println(converter(author, course))
}
func converter(author, course string) (a, c string) {
	s1 := strings.ToUpper(author)
	s2 := strings.Title(course)
	return s1, s2
}
