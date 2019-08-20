package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	name string
	age  int
}

var user map[string]interface{}

func main() {
	/*userJson2 := `
	{"username":"system",
	"password":"123456",
	"update":{
		"way1": 1,
		"way2": 2,
	}
	}
	`*/

	userJson := `{"username":"system","password":"123456","update":{"way1":1,"way2":2}}`
	json.Unmarshal([]byte(userJson), &user)

	fmt.Println(user) //打印结果:map[password:123456 username:system]
	username := user["username"]
	fmt.Println("username  ==", username) //username  == system

	update := user["update"].(map[string]interface{})
	fmt.Println(update)
	fmt.Println(update["way1"])
	update["way1"] = 100
	fmt.Println(update["way1"])

	//判断username的长度不为0,
	//if len(username) != 0 {//直接使用,报错invalid argument username (type interface {}) for len
	if len(username.(string)) != 0 { //通过.(string)转换成string类型
		fmt.Println("输入合法")
	} else {
		fmt.Println("输入不合法")
	}
}
