package main

import "fmt"

func getMax(xx ...int) int {
	max := -1
	for _, x := range xx {
		if x > max {
			max = x
		}
	}
	return max
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	res := getMax(1, 2, 3, 4)
	fmt.Println(res)
	res = getMax(x...)
	fmt.Println(res)
}
