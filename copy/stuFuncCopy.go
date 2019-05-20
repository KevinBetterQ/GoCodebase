package main

import "fmt"

func main() {
	asInputParam()
}

type Flower struct {
	name, color string
}

func asInputParam() {

	//1、结构体作为参数的用法
	f1 := Flower{"玫瑰", "红"}
	fmt.Printf("f1: %T , %v , %p , %p\n", f1, f1, &f1, &f1.name) //{玫瑰 红} , 0xc000004460 , 0xc000004460
	fmt.Println("----------------------")

	//将结构体对象作为参数。可见，也是进行的深拷贝，从指针到只想内存结构都拷贝一份。所以不会影响原结构的值
	changeInfo1(f1)                                              //{meigui 红} , 0xc0000044e0 , 0xc0000044e0
	fmt.Printf("f1: %T , %v , %p , %p\n", f1, f1, &f1, &f1.name) //{玫瑰 红} , 0xc000004460 , 0xc000004460

	fmt.Println("----------------------")

	//	将结构体指针作为参数。直接没有拷贝，就是传递的源指针和地址
	changeInfo2(&f1)                                             // &{meigui 红} , 0xc000004460 , 0xc000004460
	fmt.Printf("f1: %T , %v , %p , %p\n", f1, f1, &f1, &f1.name) //{meigui 红} , 0xc000004460 , 0xc000004460

	fmt.Println("----------------------")

}

func changeInfo1(f1 Flower) {
	f1.name = "meigui"
	fmt.Printf("f1: %T , %v , %p , %p\n", f1, f1, &f1, &f1.name)
}

func changeInfo2(f1 *Flower) {
	f1.name = "meigui"
	fmt.Printf("f1: %T , %v , %p , %p\n", f1, f1, f1, &f1.name)
}
