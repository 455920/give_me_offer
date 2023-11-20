package main

import "fmt"
import "sync"

func work(n *int, wg *sync.WaitGroup, locker *sync.Mutex, is_lock bool) {
	defer wg.Done()
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
}

func Test1() {
	var n int
	// 创建一个锁
	var locker sync.Mutex

	// 用于同步协程结束
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)

	go work(&n, &wg, &locker, false)
	go work(&n, &wg, &locker, false)
	go work(&n, &wg, &locker, false)
	go work(&n, &wg, &locker, false)

	// 等待协程结束
	wg.Wait()

	fmt.Printf("n=%d\n", n)
}


func Test2() {
	var n int
	// 创建一个锁
	var locker sync.Mutex

	// 用于同步协程结束
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)

	go work(&n, &wg, &locker, true)
	go work(&n, &wg, &locker, true)
	go work(&n, &wg, &locker, true)
	go work(&n, &wg, &locker, true)

	// 等待协程结束
	wg.Wait()

	fmt.Printf("n=%d\n", n)
}

func main() {
	// 并发累加未加锁，累加非原子操作，最终结果不符合预期
	Test1()
	// 并发累加加锁，累加为原子操作，最终结果符合预期
	Test2()
}