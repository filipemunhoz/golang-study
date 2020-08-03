package main

import "fmt"

func mainkllm() {

	styles := make([]string, 3, 10)
	fmt.Printf("Len: %d, cap: %d", len(styles), cap(styles))

	brands := []string{"Fender", "Gibson", "Jackson", "Del Vecchio", "Ibanez"}
	fmt.Printf("\nLen: %d, cap: %d\n", len(brands), cap(brands))

	newSlice := brands[2:4]
	fmt.Println(newSlice)

	mySlice := make([]int, 1, 4)
	fmt.Printf("\nlen is:%d cap is:%d\n", len(mySlice), cap(mySlice))

	for i := 0; i < 17; i++ {
		mySlice = append(mySlice, i)

		fmt.Printf("\nlen is:%d cap is:%d", len(mySlice), cap(mySlice))
	}

	newValues := []int{10, 20, 30}
	mySlice = append(mySlice, newValues...)

	fmt.Println(mySlice)

}
