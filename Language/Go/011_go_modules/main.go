package main

import "fmt"

// 导入包 根目录名 + 包相对路径
import "go_module_test/package1"
import "go_module_test/package2"

func main() {
	fmt.Printf("hello world\n")
	test1()
	Test2()
	Test3()

	file1.Test1()
	file1.Test2()

	package2.Test1()
	package2.Test2()
}
