package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	gopath := os.Getenv("GOPATH")
	fmt.Println("\nGO path:", gopath)
}
