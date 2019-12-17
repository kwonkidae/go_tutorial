package main

import "fmt"

func greet(c chan string) {
	<-c
	<-c
}

func main() {
	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()
	fmt.Println("main() started")

	c := make(chan string)

	go greet(c)
	c <- "John"

	close(c)

	c <- "Mike"
	fmt.Println("main() stopped")

}
