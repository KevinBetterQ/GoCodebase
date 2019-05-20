package main

import "fmt"

type Dog struct {
	name  string
	color string
	age   int8
	kind  string
}

func main() {
	equalCopy()
}

func equalCopy() {
	//1、struct是值类型，默认的复制就是深拷贝。即使修改也不影响原结构
	d1 := Dog{"豆豆", "黑色", 2, "二哈"}
	fmt.Printf("d1: %T , %v , %p, %p \n", d1, d1, &d1, &d1.age)
	d2 := d1 //值拷贝
	fmt.Printf("d2: %T , %v , %p, %p \n", d2, d2, &d2, &d2.age)

	d2.name = "毛毛"
	fmt.Println("d2修改后：", d2)
	fmt.Println("d1：", d1)
	fmt.Println("------------------")

	//2、直接赋值指针地址。这样修改了便会影响原结构了
	d3 := &d1
	fmt.Printf("d3: %T , %v , %p \n", d3, d3, d3)
	d3.name = "球球"
	d3.color = "白色"
	d3.kind = "萨摩耶"
	fmt.Println("d3修改后：", d3)
	fmt.Println("d1：", d1)
	fmt.Println("------------------")
}
