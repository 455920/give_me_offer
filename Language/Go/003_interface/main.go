package main

import "fmt"

type Interface1 interface {
	Printf1()
}

type Interface2 interface {
	Printf2()
}

type Interface3 interface {
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Printf1() {
	fmt.Printf("Name[%s], Age[%d]\n", p.Name, p.Age)
}

func (p Person) Printf2() {
	fmt.Printf("Name=%s, Age=%d\n", p.Name, p.Age)
}

func main() {
	person := Person{Name: "张三",
		Age: 18}
	person.Printf1()
	person.Printf2()
	var p1 Interface1
	p1 = person
	var p2 Interface2
	p2 = person
	var p3 Interface3
	// 空接口可以储存任意类型的值
	p3 = person
	p1.Printf1()
	p2.Printf2()
	// 空接口转换成其他类型，需要进行类型断言
	person_tmp, ok := p3.(Person)
	if ok {
		person_tmp.Printf1()
		person_tmp.Printf2()
	} else {
		fmt.Printf("类型断言失败")
	}
}
