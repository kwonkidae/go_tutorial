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
}

func mapForEach(arr []string, fn func(it string) int) []int {
	newArray := []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}
