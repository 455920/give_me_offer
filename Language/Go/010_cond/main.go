package main

import "sync"
import "fmt"

func main() {
	// 锁是保证资源的线程安全
	// 条件变量保证每个线程排队操作资源，确保不会出现某个线程永远抢不到资源的情况
	var lock sync.Mutex
	cond = sync.NewCond(&lock)

	// 对比
	// 有条件变量时候运行时序
	// 没有条件遍历时候运行时序
}