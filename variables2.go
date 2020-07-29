package main

import (
	"fmt"
	"os"
)

func mainoo() {

	username := os.Getenv("USERNAME")
	course := "Golang Course"

	fmt.Println("\nHi", username, " you are watching course", course)

	changeCourseWithoutPointer(course)
	fmt.Println("\nnew course no pointer=", course)

	changeCourseWithPointer(&course)
	fmt.Println("\nnew course with pointer=", course)
}

func changeCourseWithoutPointer(course string) string {

	course = "Curso Python"
	fmt.Println("\nnoPointer()=", course)

	return course
}

func changeCourseWithPointer(course *string) string {

	*course = "Curso Java"
	fmt.Println("\nwithPointer=", *course)

	return *course
}
