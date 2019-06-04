package main

// Import Package
import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()

	fmt.Println("origin : ", current.String())
	// origin :  2016-09-02 15:53:07.159994437 +0800 CST

	//StampNano
	fmt.Println("yyyy-mm-dd HH:mm:ss: ", current.Format("2006-01-02T15:04:05.781Z"))
	// yyyy-mm-dd HH:mm:ss:  2016-09-02 15:53:07.159994437
}
