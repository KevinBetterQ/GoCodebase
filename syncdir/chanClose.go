package main
/*
1、close(quit)回向所有<-quit的地方的发送一个信号
2、signal <- struct{}{}会阻塞，直到有<-signal
 */

import "fmt"

func test(quit chan int, signal chan struct{}) {
	fmt.Println("program counter came here")
	select {
	case <-quit:
		fmt.Println("pppppp")
	default:
		fmt.Println("qqqqqq")
	}
	fmt.Println("gggggg")
	<-signal
}

func test2(quit chan int) {
	select {
	case <-quit:
		fmt.Println("ttttttt") //
	default:
		fmt.Println("qqqqqq")
	}
}

func main() {
	quit := make(chan int)
	signal := make(chan struct{})
	close(quit)
	go test(quit, signal)
	test2(quit)
	signal <- struct{}{}
}
