package main

import (
	"fmt"
	"regexp"
)

func main() {
	//test()
	pouchUse()

}

func pouchUse() {
	str := `engine_daemon_engine_info{gitCommit="1.3.0-53-g938058f",kernelVersion="3.10.0-514.26.2.el7.x86_64",pouchVersion="1.3.0"} 1`
	regularStr := `^engine_daemon_engine_info{gitCommit="(.*)",kernelVersion="(.*)",pouchVersion="(.*)"} 1$`
	regular := regexp.MustCompile(regularStr)
	params := regular.FindStringSubmatch(str)
	for _, param := range params {
		fmt.Println(param)
	}
}

func test() {
	flysnowRegexp := regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	params := flysnowRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html")
	for _, param := range params {
		fmt.Println(param)
	}
}
