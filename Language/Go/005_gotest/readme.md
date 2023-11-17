在Go语言中，编写单元测试非常简单。Go的测试框架内置在标准库中，你只需要按照一定的约定编写测试函数，并使用go test命令来运行测试。

下面是一个示例，演示了如何编写和运行Go的单元测试：

创建一个名为example.go的源代码文件，其中包含一个简单的函数：
go
复制
package example

func Add(a, b int) int {
return a + b
}
创建一个名为example_test.go的测试文件，其中包含测试函数：
go
复制
package example

import "testing"

func TestAdd(t *testing.T) {
result := Add(2, 3)
expected := 5
if result != expected {
t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
}
}
在测试文件中，我们使用testing包导入了Go的测试框架，并编写了一个名为TestAdd的测试函数。在该函数中，我们调用被测试的函数Add，并使用if语句和Errorf函数来检查结果是否符合预期。

在终端中，进入包含源代码和测试文件的目录，并运行go test命令：
复制
go test
运行测试时，Go会自动查找以_test.go结尾的文件，并执行其中的测试函数。测试结果将显示在终端中。

如果测试通过，你将看到类似以下的输出：

复制
PASS
ok      example     0.001s
如果测试失败，你将看到类似以下的输出：

复制
FAIL
FAIL    example     0.001s
在失败的情况下，Go会显示详细的错误信息，指出测试失败的原因。

这就是一个简单的Go单元测试的示例。你可以根据需要编写更多的测试函数来覆盖不同的代码路径和边界情况。

希望这个示例对你有所帮助！如果你还有其他问题，请随时提问。