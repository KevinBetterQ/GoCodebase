package main

import "fmt"

type typedElem struct {
	ID   string
	Size float64
}

func main() {
	a := make([]string, 0)
	var b []string
	fmt.Println(a)
	fmt.Println(b)
	if a != nil {
		fmt.Println("not nil")
	} else if a == nil {
		fmt.Println("is nil")
	}
	if b != nil {
		fmt.Println("not nil")
	} else if b == nil {
		fmt.Println("is nil")
	}

}
