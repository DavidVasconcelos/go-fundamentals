package main

import "fmt"

func main() {
	mySlice := make([]int, 1, 4)
	fmt.Printf("Length of slice is %d and capacity is %d\n", len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Length of slice is %d and capacity is %d\n", len(mySlice), cap(mySlice))
	}

	// Length of slice is 1 and capacity is 4
	// Length of slice is 2 and capacity is 4
	// Length of slice is 3 and capacity is 4
	// Length of slice is 4 and capacity is 4
	// Length of slice is 5 and capacity is 8
	// Length of slice is 6 and capacity is 8
	// Length of slice is 7 and capacity is 8
	// Length of slice is 8 and capacity is 8
	// Length of slice is 9 and capacity is 16
	// Length of slice is 10 and capacity is 16
	// Length of slice is 11 and capacity is 16
	// Length of slice is 12 and capacity is 16
	// Length of slice is 13 and capacity is 16
	// Length of slice is 14 and capacity is 16
	// Length of slice is 15 and capacity is 16
	// Length of slice is 16 and capacity is 16
	// Length of slice is 17 and capacity is 32
}
