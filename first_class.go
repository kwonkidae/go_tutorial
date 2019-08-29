// A function can be considered as a higher-order-function only if it takes
// one or more functions as parameters or if it returns another function as a result.
package main

import "fmt"

func main() {
	list := []string{"Orange", "Apple", "Banana", "Grape"}
	out := mapForEach(list, func(it string) int {
		return len(it)
	})
	fmt.Println(out)

	add10 := add(10)
	add20 := add(20)
	add30 := add(30)

	fmt.Println(add10(5))
	fmt.Println(add20(5))
	fmt.Println(add30(5))
}

func mapForEach(arr []string, fn func(it string) int) []int {
	newArray := []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

// curring
func add(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}
