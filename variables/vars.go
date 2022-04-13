package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	global = "This var is global"
)

func main() {
	name := "David"
	course := "Getting starded with GO"
	module := "4" //Uh-oh.... that's a string!!
	clip := 2     //Needs to be integer
	//courseComplete := false
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
	// total := module + clip
	// fmt.Println("Module plus clip equals", total)

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}

	fmt.Println("memory address of *course* variable is ", &course)

	var pointer *string = &course
	fmt.Println("Pointing course variable at address", pointer, "which holds this value,", *pointer)

	fmt.Println("\nHi", name, "your current course is", course)
	updateCourseByValue(course)
	fmt.Println("You're currently watching", course)

	updateCourseByReference(&course)
	fmt.Println("You're currently watching", course)
}

func updateCourseByValue(course string) string {
	course = "Getting started with Docker"
	fmt.Println("Updated course to", course)
	return course
}

func updateCourseByReference(course *string) string {
	*course = "Getting started with Kubernetes"
	fmt.Println("Updated course to", *course)
	return *course
}
