package main

import (
	"fmt"
)

func main() {
	
	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	
	// updated
	testMap["A"] = 100
	
	// Create
	testMap["F"] = 1973
	testMap["G"] = 53
	
	fmt.Println(testMap)
	
	delete(testMap, "F")
	fmt.Println(testMap)
}
