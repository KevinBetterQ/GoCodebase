package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	mainStart_end_ser := time.Now()
	n := 30
	//n := 100
	//n := 1000
	rsyncGet(n)
	//Get(n)
	mainEnd := time.Since(mainStart_end_ser)
	fmt.Println("main cost", mainEnd.String())
}

func Get(n int) {
	for i := 0; i < n; i++ {
		get()
	}
}

func rsyncGet(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			get()
		}()
	}
	wg.Wait()
}

func get() {
	getStart := time.Now()
	//fmt.Println("start: ", getStart)
	//http.Get("https://i.sigma.alibaba-inc.com/store/api/v1/label/job/values")
	resp, err := http.Get("https://i.sigma.alibaba-inc.com/store/daily/api/v1/query?query={__name__=~%22graph_usage_avail|graph_usage_percentage%22,%20instance=%2211.12.179.220%22}")
	//resp, err := http.Get("http://i.sigma.alibaba.net/store/api/v1/label/job/values")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	getEnd := time.Since(getStart)
	fmt.Println("get cost", getEnd.String())
}
