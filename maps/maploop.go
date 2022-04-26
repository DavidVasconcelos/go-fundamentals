package main

import "fmt"

func main() {
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
	}

	for mapKey, mapValue := range testMap {
		fmt.Printf("Key is: %v Value is: %v\n", mapKey, mapValue)
	}

	// Key is: B Value is: 2 Go doesn't guarantee iteration order with range loops
	// Key is: F Value is: 6
	// Key is: I Value is: 9
	// Key is: A Value is: 1
	// Key is: C Value is: 3
	// Key is: D Value is: 4
	// Key is: E Value is: 5
	// Key is: G Value is: 7
	// Key is: H Value is: 8

	testMap["A"] = 100
	testMap["J"] = 1973
	fmt.Println(testMap)
	// map[A:100 B:2 C:3 D:4 E:5 F:6 G:7 H:8 I:9 J:1973] order by key

	delete(testMap, "J")
	fmt.Println(testMap)
	// map[A:100 B:2 C:3 D:4 E:5 F:6 G:7 H:8 I:9]
}
