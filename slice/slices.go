package main

import "fmt"

func main() {
	// courses := make([]string, 5, 10) //type, length, expected size
	// courses[0] = "Docker & Kubernetes: The Big Picture"
	// courses[1] = "Getting started with Docker"
	// courses[2] = "Getting started with Kubernetes"

	courses := []string{"Docker & Kubernetes: The Big Picture",
		"Getting started with Docker", "Getting started with Kubernetes"}

	fmt.Printf("Length of slice is %d and capacity is %d\n", len(courses), cap(courses))

	for _, course := range courses {
		fmt.Println(course)
	}

}
