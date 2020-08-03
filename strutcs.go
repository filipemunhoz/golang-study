package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("//tmp//file.txt")

	if err != nil {
		fmt.Println("Could not open the file", err)
	}

	type person struct {
		Name  string
		Email string
		Age   int
	}

	P := person{
		Name:  "Filipe",
		Email: "test@test.com",
		Age:   36,
	}

	fmt.Println(P)
	fmt.Println(P.Email)
}
