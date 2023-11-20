package main

import "fmt"

func Test1() {
	// k为string, v为int
	kv := make(map[string]int)
	// 添加一个元素
	kv["key1"] = 1
	kv["key2"] = 2
	kv["key3"] = 3
	kv["key4"] = 4

	// 查找元素
	value1, ok1 := kv["key1"]
	if ok1 {
		fmt.Printf("find key1=%d\n", value1)
	} else {
		fmt.Printf("not find key1\n")
	}

	value5, ok5 := kv["key5"]
	if ok5 {
		fmt.Printf("find key5=%d\n", value5)
	} else {
		fmt.Printf("not find key5\n")
	}

	// 删除元素
	delete(kv, "key1")
	// 查找元素
	value1, ok1 = kv["key1"]
	if ok1 {
		fmt.Printf("find key1=%d\n", value1)
	} else {
		fmt.Printf("not find key1\n")
	}

	// 遍历元素
	for key, value := range kv {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func main() {
	Test1()
}
