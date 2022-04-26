package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(mySlice[4])

	mySlice[1] = 0
	fmt.Println(mySlice)

	sliceOfSlice := mySlice[2:5] //include 2 and exclude 5
	fmt.Println(sliceOfSlice)

	sliceOfSliceIndexZero := mySlice[:5] //include 0 and exclude 5
	fmt.Println(sliceOfSliceIndexZero)

	sliceOfSliceEndImplice := mySlice[5:] //include 0 and include 10
	fmt.Println(sliceOfSliceEndImplice)
}
