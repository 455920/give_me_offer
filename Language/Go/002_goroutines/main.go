package main

import "fmt"
import "time"

func work1() {
	for i := 0; i < 2; i++ {
		fmt.Printf("work1:%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func work2() {
	for i := 0; i < 4; i++ {
		fmt.Printf("work2:%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func work3() {
	for i := 0; i < 10; i++ {
		fmt.Printf("work3:%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go work1()
	go work2()
	go work3()
	time.Sleep(8 * time.Second)
}
