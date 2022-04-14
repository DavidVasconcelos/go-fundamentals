package main

import "fmt"

func main() {
	bestFinish := championshipFinishes(12, 5, 6, 4, 3, 3)

	fmt.Println(bestFinish)
}

func championshipFinishes(finishes ...int) int {
	best := finishes[0]
	for _, finish := range finishes {
		if finish < best {
			best = finish
		}
	}
	return best
}
