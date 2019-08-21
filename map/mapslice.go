package main

import "fmt"

type Data struct {
	name string
}

func main() {
	m := map[string][]Data{"x": {{"one"}, {"two"}}}
	fmt.Println(m["x"]) //prints: &{four}
	data2 := Data{"one2"}
	m["x"] = append(m["x"], data2)
	fmt.Println(m["x"]) //prints: &{four}
	m["x"][0].name = "three"
	fmt.Println(m["x"]) //prints: &{four}

}
