package main

import "fmt"

var funcVar func(int) int

func myFunc(a int) int {
	return a + 2
}

func main() {
	funcVar = myFunc
	fmt.Println(funcVar(3))
}
