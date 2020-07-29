package main

import (
	"fmt"
	"os"
)

func mainxx() {

	_, err := os.Open("//tmp//file.txt")

	if err != nil {
		fmt.Println("Could not open the file", err)
	}

}
