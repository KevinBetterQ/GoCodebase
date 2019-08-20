package main

import (
	"fmt"
	"regexp"
)

func main() {
	//test()
	versionMetrics := `sadasengine_daemon_engine_info{commit="commit866",version="3.10.0-514.26.2.el7.x86_64",kernel="1.3.0"} 1`
	regularCommit := `^.*commit="(.*?)".*$`
	regular := regexp.MustCompile(regularCommit)
	params := regular.FindStringSubmatch(versionMetrics)
	fmt.Println(params[1])

	regularVersion := `^.*version="(.*?)".*$`
	regular = regexp.MustCompile(regularVersion)
	params = regular.FindStringSubmatch(versionMetrics)
	fmt.Println(params[1])

	regularKernel := `^.*kernel="(.*?)".*$`
	regular = regexp.MustCompile(regularKernel)
	params = regular.FindStringSubmatch(versionMetrics)
	fmt.Println(params[1])

}

func pouchUse() {
	str := `sadasengine_daemon_engine_info{commit="",version="3.10.0-514.26.2.el7.x86_64",kernel="1.3.0"} 1`
	regularStr := `^.*commit="(.*)",version="(.*)",kernel="(.*)".*$`
	regular := regexp.MustCompile(regularStr)
	params := regular.FindStringSubmatch(str)
	fmt.Println("0: " + params[0])
	fmt.Println("1: " + params[1])
	fmt.Println("2: " + params[2])
	fmt.Println("3: " + params[3])
	/*for _, param := range params {
		fmt.Println(param)
	}*/
}

func test() {
	flysnowRegexp := regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	params := flysnowRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html")
	for _, param := range params {
		fmt.Println(param)
	}
}
