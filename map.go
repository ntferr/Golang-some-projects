package main

import "fmt"

func main() {
	m := map[string]int{
		"Nathan": 25,
		"Ted":    30,
	}
	
	v, ok := m["Tod"]
	fmt.Println(v, ok)
	
	if v, ok := m["Nathan"]; ok {
		fmt.Println("Exist.", v, ok)
	}
	
	for k, v := range m {
		fmt.Println(k, v)
	}
}
