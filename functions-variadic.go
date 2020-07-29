package main

import (
	"fmt"
)

func mainsss() {

	highest := getHighest(12, 5, 89, 45, 23, 4, 9, 11)

	fmt.Println("Highest is:", highest)

}

func getHighest(values ...int) int {

	highest := values[0]

	for _, i := range values {
		if i > highest {
			highest = i
		}
	}
	return highest
}
