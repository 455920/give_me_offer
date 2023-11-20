package main

import "fmt"
import "sync"

func work(n *int, endCh chan<- int, locker *sync.Mutex, is_lock bool) {
	for i:=0; i< 1000000; i++ {
		if (is_lock) {
			locker.Lock()
		}
		// 解引用
		(*n)++

		if (is_lock) {
			locker.Unlock()
		}
	}
	close(endCh)
}

func Test1() {
	var n int
	// 创建一个锁
	var locker sync.Mutex

	// 用于同步协程结束
	endCh1 := make(chan int)
	endCh2 := make(chan int)
	endCh3 := make(chan int)
	endCh4 := make(chan int)
	go work(&n, endCh1, &locker, false)
	go work(&n, endCh2, &locker, false)
	go work(&n, endCh3, &locker, false)
	go work(&n, endCh4, &locker, false)

	// 等待协程结束
	_ = <-endCh1
	_ = <-endCh2
	_ = <-endCh3
	_ = <-endCh4

	fmt.Printf("n=%d\n", n)
}


func Test2() {
	var n int
	// 创建一个锁
	var locker sync.Mutex

	// 用于同步协程结束
	endCh1 := make(chan int)
	endCh2 := make(chan int)
	endCh3 := make(chan int)
	endCh4 := make(chan int)
	go work(&n, endCh1, &locker, true)
	go work(&n, endCh2, &locker, true)
	go work(&n, endCh3, &locker, true)
	go work(&n, endCh4, &locker, true)

	// 等待协程结束
	_ = <-endCh1
	_ = <-endCh2
	_ = <-endCh3
	_ = <-endCh4

	fmt.Printf("n=%d\n", n)
}

func main() {
	// 并发累加未加锁，累加非原子操作，最终结果不符合预期
	Test1()
	// 并发累加加锁，累加为原子操作，最终结果符合预期
	Test2()
}