package main

import "fmt"

func main() {
	m := map[string]int{
		"Nathan": 25,
		"Ted":    30,
	}
	
	v, ok := m["Tod"]
	fmt.Println(v, ok)
	
	fmt.Println(m)
	
	if v, ok := m["Nathan"]; ok {
		fmt.Println("Exist.", v, ok)
	}
	
	m["Jess"] = 29
	
	fmt.Println(m)
	
	for k, v := range m {
		fmt.Println(k, v)
	}
}
