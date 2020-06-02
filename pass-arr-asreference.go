package main

import (
	"fmt"
)

func foo(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}

func main() {
	x := [3]int{1, 2, 3}
	foo(&x)
	for _, a := range x {
		fmt.Println(a)
	}
}
