package main

import "fmt"

func firstFunc(fn func(int) int, i int) int {
	return fn(i)
}

func myFunc2(a int) int {
	return a + 2
}

func main() {

	// function as argument
	res := firstFunc(myFunc2, 2)
	fmt.Println(res)

	// anonymous function
	res = firstFunc(func(a int) int { return a + 2 }, 3)
	fmt.Println(res)

}
