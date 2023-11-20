package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	// 逐行读取输入数据
	for scanner.Scan() {
		// 获取输入的文本
		text := scanner.Text()

		// 打印输入的文本
		fmt.Println("Input:", text)

		// 如果输入的是"exit"，则退出循环
		if text == "exit" {
			break
		}

		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Printf("err: %s", err.Error())
			break
		}
	}

	// 检查Scanner是否出现错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
	}
}
