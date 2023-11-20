package main

import "fmt"

func test1() {
	fmt.Printf("example1-test1\n")
}

func Test2() {
	fmt.Printf("example1-Test2\n")
}

func Test3() {
	test1()
}

func Test4() {
	test1()
}
