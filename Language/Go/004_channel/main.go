package main

import (
	"fmt"
	"time"
)

func work1(ch_in <-chan int, ch_out chan<- int) {
	// 读取到为nil时 表示channel已经关闭数据已经读完
	// for语句循环读取
	for i:= range ch_in {
		fmt.Printf("work1_read\n")
		ch_out<-i*i
		fmt.Printf("work1_push\n")
	}
	close(ch_out)
	fmt.Printf("work1_end")
}
func work2(ch_in <-chan int) {
	// 读取到为nil时 表示channel已经关闭数据已经读完
	for true {
		if val, ok := <-ch_in; ok {
			fmt.Printf("work2_read i=%d\n", val)
			time.Sleep(1 * time.Second)
			} else {
			break;
		}
	}
	fmt.Printf("work2_end\n")
}

func main() {
	// 有缓存和无缓存运行的时序不同
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 2)

	go work1(ch1, ch2)
	go work2(ch2)

	for i := 0; i < 10; i++ {
		fmt.Printf("main_push\n")
		ch1<-i
	}
	fmt.Printf("main_end\n")
	close(ch1)
	time.Sleep(12 * time.Second)
}
