package main

import "fmt"

func main() {
	defer fmt.Println("this is fun deferred")
	fmt.Println("fun immediate")
}
