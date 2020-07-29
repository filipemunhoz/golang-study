package main

import (
	"fmt"
	"math/rand"
	"time"
)

func maink() {

	if age, size := 36, 100; age > 10 && size > 20 {

		fmt.Println("Yes, who said this is an elegant code is completely crazy")
	}

	types := "flex"

	// Example 1
	switch types {
	case "flex":
		fmt.Println("Yes, this is flex")
	case "none":
		fmt.Println("none selected")
	default:
		fmt.Println("default option")
	}

	// Example 2
	switch number := random(); number {
	case 0, 2, 4, 6, 8:
		fmt.Println("Par:", number)
	case 1, 3, 5, 7, 9:
		fmt.Println("Impar", number)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
