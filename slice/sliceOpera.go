package main

import (
	"fmt"
	"reflect"
)

func main() {
	warnings := make([]string, 0, 64)
	fmt.Println(warnings)
	fmt.Println(len(warnings))
	expect := make([]string, 0, 64)
	if warnings == nil {
		fmt.Println("is nil")
	} else if reflect.DeepEqual(warnings, expect) {
		fmt.Println("equal expect")
	} else {
		fmt.Println("not nil")
	}
}
