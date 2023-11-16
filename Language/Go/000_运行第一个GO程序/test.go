// 在Go语言中，程序的入口点是一个名为 main 的函数。当你运行一个Go程序时，程序会自动执行 main 函数
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

// 在上述示例中，我们定义了一个 main 函数，它没有任何参数和返回值。在 main 函数中，我们使用 fmt.Println 函数打印了一条简单的消息。
//
// 当你运行这个程序时，它会输出 Hello, World!。这是因为 main 函数是程序的入口点，程序会从这里开始执行。
//
// 需要注意的是，一个Go程序只能有一个 main 函数，并且它必须位于 main 包中。在程序中，你可以在 main 函数中调用其他函数、创建变量、导入包等。
//
// 除了 main 函数，Go语言还支持在包级别定义的 init 函数。init 函数在程序启动时自动执行，用于初始化包级别的变量和执行一些初始化操作。但是，init 函数不是程序的入口点，它会在 main 函数执行之前被调用。
//
// 总结来说，Go程序的入口点是 main 函数，它是程序的起点，从这里开始执行。