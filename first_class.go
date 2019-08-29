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

var holder = map[string]int{}

func sum(a, b int) int {
	c := a + b
	holder[fmt.Sprintf("%d+%d", a, b)] = c
	return c
}

func normalFactorial(num int) int {
	result := 1
	for ; num > 0; num-- {
		result *= num
	}
	return result
}

func factorialRecursion(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorialRecursion(num-1)
}

func factorialTailRec(num int) int {
	return factorial(1, num)
}

func factorial(accumulator, val int) int {
	if val == 1 {
		return accumulator
	}
	return factorial(accumulator*val, val-1)
}
