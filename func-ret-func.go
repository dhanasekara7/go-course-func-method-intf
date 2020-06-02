package main

import "fmt"

func func_ret_func(aOrbFunc bool) func(int) int {
	if aOrbFunc {
		return func(x int) int { return x + 2 }
	} else {
		return func(x int) int { return x + 3 }
	}
}
func main() {
	res := func_ret_func(true)(2)
	fmt.Println(res)
	res = func_ret_func(false)(2)
	fmt.Println(res)
}
