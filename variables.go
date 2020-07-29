package main

import (
	"fmt"
	"reflect"
)

var (
	name   = "Filipe"
	course = "Goland"
	module = 3.2
	city   string
)

func mainvss() {

	surname := "ALoha"
	ptr := &module

	fmt.Println("Name = ", name, " and its type of", reflect.TypeOf(name))
	fmt.Println("Module = ", module, " and its type of", reflect.TypeOf(module))
	fmt.Println("status:", surname)
	fmt.Println("Module variable address is:", ptr, "and the value is", *ptr)
}
