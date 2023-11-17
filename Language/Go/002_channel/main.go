package main

import (
	"fmt"
)

func Test1() {
	fmt.Println("==================== Test1 =======================")
	// 创建一个切片 引用的 底层数组 数据长度为10，最大容量为10
	s1 := make([]int, 10, 10)
	// 初始化s1
	for i := 0; i < 10; i++ {
		s1[i] = i
	}
	fmt.Printf("s1-Length:%d\n", len(s1))
	fmt.Printf("s1-Capacity:%d\n", cap(s1))

	// s1和s2引用同一块底层数组, 打印s1和s2是相同的
	s2 := s1
	fmt.Printf("s2-Length:%d\n", len(s2))
	fmt.Printf("s2-Capacity:%d\n", cap(s2))
	fmt.Printf("s1-Address:%p\n", &s1[0])
	fmt.Printf("s2-Address:%p\n", &s2[0])
	fmt.Printf("s1[1]:%d\n", s1[1])
	fmt.Printf("s2[1]:%d\n", s2[1])

	// 修改s1, s2也会一起变, 因为引用的同一个底层数组
	s1[1] = 123
	fmt.Printf("s1[1]:%d\n", s1[1])
	fmt.Printf("s2[1]:%d\n", s2[1])

	// s1增加一个数据,因为容量不够,s1重新引用到一个新的更大底层数组,旧底层数组的数据会拷贝的新底层数组
	// s1和s2指向的不再是一个底层数组
	s1 = append(s1, 999)
	fmt.Printf("s1-Address:%p\n", &s1[0])
	fmt.Printf("s2-Address:%p\n", &s2[0])
	// 修改s1后, s2不会一起变了, 因为引用的不是同一个底层数组了
	s1[1] = 321
	fmt.Printf("s1[1]:%d\n", s1[1])
	fmt.Printf("s2[1]:%d\n", s2[1])
}

func Test2() {
	fmt.Println("==================== Test2 =======================")
	// 创建一个切片 引用的 底层数组 数据长度为10，最大容量为11
	s1 := make([]int, 10, 11)
	// 初始化s1
	for i := 0; i < 10; i++ {
		s1[i] = i
	}
	fmt.Printf("s1-Length:%d\n", len(s1))
	fmt.Printf("s1-Capacity:%d\n", cap(s1))

	// s1和s2引用同一块底层数组, 打印s1和s2是相同的
	s2 := s1
	fmt.Printf("s2-Length:%d\n", len(s2))
	fmt.Printf("s2-Capacity:%d\n", cap(s2))
	fmt.Printf("s1-Address:%p\n", &s1[0])
	fmt.Printf("s2-Address:%p\n", &s2[0])
	fmt.Printf("s1[1]:%d\n", s1[1])
	fmt.Printf("s2[1]:%d\n", s2[1])

	// 修改s1, s2也会一起变, 因为引用的同一个底层数组
	s1[1] = 123
	fmt.Printf("s1[1]:%d\n", s1[1])
	fmt.Printf("s2[1]:%d\n", s2[1])

	// s1增加一个数据,容量足够,无需另外扩容
	// s1和s2指向一个底层数组
	s1 = append(s1, 999)
	fmt.Printf("s1-Address:%p\n", &s1[0])
	fmt.Printf("s2-Address:%p\n", &s2[0])
	// 修改s1后, s2不会一起变了, 因为引用的不是同一个底层数组了
	s1[1] = 321
	fmt.Printf("s1[1]:%d\n", s1[1])
	fmt.Printf("s2[1]:%d\n", s2[1])
}

func main() {
	Test1()
	Test2()
}
