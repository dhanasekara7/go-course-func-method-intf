package main

import (
	"fmt"
)

func foo(x1 []int) {
	x1[0] = x1[0] + 1
}

func main() {
	x := []int{1, 2, 3}
	foo(x)
	for _, a := range x {
		fmt.Println(a)
	}
}
