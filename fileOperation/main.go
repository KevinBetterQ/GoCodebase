package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

// 参考：https://colobu.com/2016/10/12/go-file-operations/
func main() {
	fileCreateWrite()
}

func fileCreateWrite() {
	file, err := os.Create("fileOperation/createFile.txt")
	if err != nil {
		logrus.Errorln(err)
	}
	defer file.Close()
	byteSlice := []byte("test create file!\n")
	n, err := file.Write(byteSlice)
	if err != nil {
		logrus.Errorln(err)
	}
	file.Write(byteSlice)
	fmt.Println("Write bytes: ", n)
}

func fileWrite() {
	file, err := os.OpenFile(
		"fileOperation/test.txt",
		os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		logrus.Errorln(err)
	}
	defer file.Close()
	// 写字节到文件
	byteSlice := []byte("Bytessss!\n")
	byteSliceWritten, err := file.Write(byteSlice)
	if err != nil {
		logrus.Errorln(err)
	}
	fmt.Printf("Wrote %d bytes.\n", byteSliceWritten)
}
