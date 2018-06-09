package main

import "fmt"

func main()  {
	fmt.Println(add(1, 2))
	fmt.Println(new_add(1, 2))
}

func add(x int, y int) int {
	return x + y
}

func new_add(x, y int) int {
	return x + y
}
