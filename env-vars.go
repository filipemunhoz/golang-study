package main

import (
	"fmt"
	"os"
)

func maincc() {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
