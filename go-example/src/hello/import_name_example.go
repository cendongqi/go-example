package main

import (
	"fmt"
	"math"
)

func main() {
	// 只能使用已导出的变量或方法，已导出的变量名和方法名首字母大写，小写的只能在本包内使用
	// 首字母大写类似于Java中的public修饰符，小写相当于private

	//fmt.Println(math.pi)
	fmt.Println(math.Pi)
}
