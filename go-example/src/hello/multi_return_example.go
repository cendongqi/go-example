package main

import "fmt"

func main() {
	a, b, c := swap("hello", "world", "go")
	fmt.Println(a, b, c)
}

func swap(x, y, z string) (string, string, string)  {
	return x, y, z
}
