package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "David"
	course := "getting starded with GO"

	fmt.Println(converter(author, course))
}

func converter(author, course string) (a, c string) {
	author = strings.ToUpper(author)
	course = strings.Title(course)

	return author, course
}
