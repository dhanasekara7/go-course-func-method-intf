package main

import "fmt"

type MyInt int

func (m MyInt) DoubleIt() int {
	return int(m * 2)
}

func main() {
	v := MyInt(3)
	res := v.DoubleIt()
	fmt.Println(res)
}
