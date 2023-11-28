package main

import "fmt"

func defer_test1() {
	fmt.Printf("defer_test1\n")
}

func defer_test2() {
	fmt.Printf("defer_test2\n")
}

func defer_test3() {
	fmt.Printf("defer_test3\n")
}

func main() {
	defer defer_test1()
	defer defer_test2()
	defer defer_test3()
	// defer执行顺序, 先进后出
	// 执行顺序defer_test3 -> defer_test2 -> defer_test1
}