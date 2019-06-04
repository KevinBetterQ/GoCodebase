package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Add(-time.Hour))
	fmt.Println(time.Now().Add(-time.Hour * 24))
	fmt.Println(time.Now().Add(-time.Hour * 25))

}
