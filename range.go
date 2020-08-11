package main

import (
	"fmt"
)

func main() {

	styles := []string{"Rock", "Country", "Blues", "Jazz"}

	printStyles(styles)
}

func printStyles(styles []string) {

	for i := range styles {
		fmt.Println("Style:", styles[i])
	}

	for i, style := range styles {
		fmt.Printf("Style: %v at index %v\n", style, i)
	}

	for _, style := range styles {
		fmt.Printf("Style: %v\n", style)
	}
}
