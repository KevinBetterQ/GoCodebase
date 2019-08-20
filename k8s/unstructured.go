package main

import (
	"fmt"
	"strings"
)

func main() {
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	mm := make(map[string]interface{})
	mm["country"] = countryCapitalMap

	var obj interface{} = mm

	demo := map[string]interface{}{
		"kind":       "Pod",
		"apiVersion": "v1",
		"metadata":   map[string]interface{}{"name": "test"},
	}
	str := "metadata.name"
	arrStr := strings.Split(str, ".")

	fmt.Println("开始使用======")
	fmt.Println(obj)
	fmt.Println("获取测试======")
	fmt.Println("demo:", demo)
	arr := []string{"metadata", "name"}
	fmt.Println(NestedFieldNoCopy(demo, arrStr...))
	fmt.Println("修改测试======")
	err := setNestedFieldNoCopy(demo, "test3", arr...)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(demo)

}

func setNestedFieldNoCopy(obj map[string]interface{}, value interface{}, fields ...string) error {
	m := obj

	for i, field := range fields[:len(fields)-1] {
		if val, ok := m[field]; ok {
			if valMap, ok := val.(map[string]interface{}); ok {
				m = valMap
			} else {
				return fmt.Errorf("value cannot be set because %v is not a map[string]interface{}", fields[:i+1])
			}
		} else {
			newVal := make(map[string]interface{})
			m[field] = newVal
			m = newVal
		}
	}
	m[fields[len(fields)-1]] = value
	return nil
}

func NestedFieldNoCopy(obj map[string]interface{}, fields ...string) (interface{}, bool, error) {
	var val interface{} = obj

	for i, field := range fields {
		if m, ok := val.(map[string]interface{}); ok {
			val, ok = m[field]
			if !ok {
				return nil, false, nil
			}
		} else {
			return nil, false, fmt.Errorf("failed %v", i)
		}
	}
	return val, true, nil
}
