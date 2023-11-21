在Go语言中，可以使用net/http包来创建和运行HTTP服务。以下是一个简单的示例，演示如何在Go中创建一个基本的HTTP服务：
```
package main

import (
"fmt"
"net/http"
)

func main() {
// 定义处理函数
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Hello, World!")
})

	// 启动HTTP服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server error:", err)
	}
}
```
在上述代码中，我们首先定义了一个处理函数，它会在根路径（"/"）上返回一个简单的"Hello, World!"消息。然后，我们使用http.HandleFunc函数将处理函数与根路径绑定。

最后，我们使用http.ListenAndServe函数启动HTTP服务，监听在本地的8080端口。如果启动过程中出现错误，我们会打印错误信息。

你可以在浏览器中访问http://localhost:8080，应该能够看到"Hello, World!"消息。

这只是一个简单的示例，你可以根据需要添加更多的处理函数和路由规则。net/http包提供了丰富的功能，例如处理静态文件、路由器、中间件等。