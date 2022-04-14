package main

import "fmt"

func main() {
	// dddLengthMins := 60  // Docker deep dive course
	// cawLengthMins := 120 // Containers on AWS Wavelength course

	if dddLengthMins, cawLengthMins := 60, 120; dddLengthMins > cawLengthMins {
		fmt.Println("Docker deep dive is longer than Containers on AWS Wavelength")
	} else if dddLengthMins == cawLengthMins {
		fmt.Println("Both courses are the same length")
	} else {
		fmt.Println("Containers on AWS Wavelength must be longer than Containers on AWS")
	}

}
