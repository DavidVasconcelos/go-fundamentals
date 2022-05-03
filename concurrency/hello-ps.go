package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// func() { //anonymous function
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Println("Hello")
	// }() //self executing

	// func() {
	// 	fmt.Println("Pluralsight")
	// }()

	//now with go routines

	var waitGrp sync.WaitGroup
	waitGrp.Add(2) //one for each go routine, this way main's going to wait until both report

	go func() { //go routine
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Pluralsight")
	}()
	waitGrp.Wait()

	// myChann := make(chan int, 5) Type and Buffers, if buffered channel is full, then, for sure, any Go routines wanting to use it are going to block until it frees up
}
