package main

import (
	"fmt"
	"time"
)

// 参考：http://blog.studygolang.com/2014/02/go_crontab/
//https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter04/04.4.md
func main() {
	input := make(chan interface{})
	go func() {
		for i := 0; i < 5; i++ {
			input <- i
		}
		input <- "hello input"
	}()
	// 值定时一次，可以辅助使用 t1.Reset(time.Second * 3) 实现和NewTicker相同效果
	t1 := time.NewTimer(time.Second * 3)
	// 一直定时执行
	t2 := time.NewTicker(time.Second * 5)

	for {
		select {
		case msg := <-input:
			fmt.Println(msg)
		case <-t1.C:
			fmt.Println("3s timer")
			t1.Reset(time.Second * 3)
		case <-t2.C:
			fmt.Println("5s ticker")
		}
	}
}
