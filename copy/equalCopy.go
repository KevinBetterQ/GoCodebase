package main

import "fmt"

func main() {
	sliceCppy()
}

/*
slice的 = 或者 函数参数参数传递时：
由于slice整体包含 指针和所指向的内存区域，
拷贝或传递时产生了新的slice，是将指针进行值传递了，但新的slice里面的指针所指向的内存区域是没变的
*/
func sliceCppy() {
	arr := [5]int{1, 3, 5, 6, 7}
	fmt.Printf("addr:%p\n", &arr)
	// addr:0xc00005e060
	s1 := arr[:]
	fmt.Printf("addr:%p\n", &s1)
	// addr:0xc000044420
	s3 := s1
	fmt.Printf("addr:%p\n", &s3)
	// addr:0xc000044440 和复制的s1的地址不同
	fmt.Printf("addr:%p\n", &s3[0])
	// addr:0xc00005e060 和最初的数组地址相同
}
