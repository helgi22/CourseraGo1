package main

import "fmt"
import "encoding/json"

func main() {
	var name string
	var addr string
	m := make(map[string]string)
	fmt.Printf("Enter name")
	fmt.Scan(&name)
	fmt.Printf("Enter address")
	fmt.Scan(&addr)
	m["name"] = name
	m["addr"] = addr
	fmt.Println(m)
	mjson, err := json.Marshal(m)
	if err == nil {
		fmt.Printf("JSON object: ")
		fmt.Println(string(mjson))
	}
	err1 := json.Unmarshal(mjson, &m)
	if err1 == nil {
		fmt.Print(m)
	}
}
