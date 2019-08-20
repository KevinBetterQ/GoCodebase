package main

import "fmt"

type typedElem struct {
	ID   int
	Size string
}

func main() {
	var demo typedElem
	demo.ID = 1
	if demo.Size == "" {
		fmt.Println(demo)
	}

}
