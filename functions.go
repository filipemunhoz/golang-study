package main

import (
	"fmt"
	"strings"
)

func mains() {

	name := "filipe munhoz"
	course := "go lange course online plataform"

	fmt.Println(converter(name, course))
}

func converter(s1, s2 string) (a, b string) {

	s1 = strings.ToUpper(s1)
	s2 = strings.Title(s2)

	return s1, s2
}
