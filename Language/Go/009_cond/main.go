package main

import (
	"fmt"
	"sync"
	"time"
)

func consumer(workName string, i *int, sum *int, cond *sync.Cond, mutex *sync.Mutex, isUseCond bool) {
	for true {
		time.Sleep(time.Millisecond * 1000)
		mutex.Lock()
		if *i <= 0 {
			// 释放锁和阻塞
			if isUseCond {
				cond.Wait()
			} else {
				continue
			}
		}
		*i--
		*sum++
		fmt.Printf("%s, sum=%d\n", workName, *sum)
		mutex.Unlock()
	}
}

func producer(workName string, i *int, cond *sync.Cond, mutex *sync.Mutex, isUseCond bool) {
	for true {
		mutex.Lock()
		if *i < 20 {
			*i++
			fmt.Printf("%s, i=%d\n", workName, *i)
		}
		if isUseCond {
			cond.Signal()
		}
		mutex.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// 锁是保证资源的线程安全, 避免同时操作资源导致资源出错
	// 条件变量保证每个线程排队操作资源，确保不会出现某个线程永远抢不到资源的情况
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	// 对比 多消费者 单生产者
	// 有条件变量时候运行时序, 生产者是一定可以拿到锁的
	var i int
	var sum int
	go consumer("consumerA1", &i, &sum, cond, &mutex, true)
	go consumer("consumerA2", &i, &sum, cond, &mutex, true)
	go consumer("consumerA3", &i, &sum, cond, &mutex, true)
	go consumer("consumerA4", &i, &sum, cond, &mutex, true)
	go consumer("consumerA5", &i, &sum, cond, &mutex, true)
	go consumer("consumerA6", &i, &sum, cond, &mutex, true)
	go consumer("consumerA7", &i, &sum, cond, &mutex, true)
	go consumer("consumerA8", &i, &sum, cond, &mutex, true)
	go consumer("consumerA9", &i, &sum, cond, &mutex, true)
	go consumer("consumerA10", &i, &sum, cond, &mutex, true)
	go producer("producerA1", &i, cond, &mutex, true)

	// 没有条件遍历时候运行时序, 存在生产者抢不到锁的情况
	go consumer("consumerB1", &i, &sum, cond, &mutex, false)
	go consumer("consumerB2", &i, &sum, cond, &mutex, false)
	go consumer("consumerB3", &i, &sum, cond, &mutex, false)
	go consumer("consumerB4", &i, &sum, cond, &mutex, false)
	go consumer("consumerB5", &i, &sum, cond, &mutex, false)
	go consumer("consumerB6", &i, &sum, cond, &mutex, false)
	go consumer("consumerB7", &i, &sum, cond, &mutex, false)
	go consumer("consumerB8", &i, &sum, cond, &mutex, false)
	go consumer("consumerB9", &i, &sum, cond, &mutex, false)
	go consumer("consumerB10", &i, &sum, cond, &mutex, false)
	go producer("producerB1", &i, cond, &mutex, false)

	// 死循环等待
	for true {
	}
}
