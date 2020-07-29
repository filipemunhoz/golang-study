package main

import (
	"fmt"
	"reflect"
)

func mainv3(){

	a := 10.0000000000
	b := 3

	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is type of:", reflect.TypeOf(b))

	c := int(a) + b

	fmt.Println("\nC has value", c, "and C is type of:", reflect.TypeOf(c))
}