package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
	<-c
}
func main() {
	// var c chan int
	// fmt.Println(c)
	fmt.Println("main() started")
	c := make(chan string)
	done := make(chan string)

	go greet(c)
	go func() {
		c <- "John"
		done <- "end"
	}()
	<-done
	fmt.Println("main() stopped")

	// fmt.Printf("type of `c` is %T\n", c)
	// fmt.Printf("value of `c` is %v\n", c)

	a := make([]int, 10)
	fmt.Printf("value %v\n", a)
}
