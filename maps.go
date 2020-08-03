package main

import "fmt"

func mainop() {

	// Example 1
	persons := make(map[string]int)

	persons["Filipe"] = 36
	persons["Bob"] = 55
	persons["Bruce"] = 63

	newPersons := map[string]int{

		"Tarja": 40,
		"Floor": 40,
	}
	fmt.Printf("Persons: \t%v\nNew Persons: \t%v\n", persons, newPersons)

	// Example 2
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	for key, value := range testMap {
		fmt.Printf("Key: %v, Value: %d\n", key, value)
	}

	// Example 3
	otherMap := map[string]int{"A": 100, "B": 200, "C": 300, "D": 400, "E": 500}
	otherMap["A"] = 666
	otherMap["F"] = 999
	delete(otherMap, "B")
	fmt.Println(otherMap)
}
