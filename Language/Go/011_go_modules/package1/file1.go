package file1

// 在Go语言中，一个目录下只能定义一个包。这是Go语言的约定和规范。

import "fmt"

func test1() {
	fmt.Printf("package1-test\n")
}

func Test1() {
	fmt.Printf("package1-Test1\n")
}

func Test2() {
	fmt.Printf("package1-Test2\n")
}
