package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

var fn string
var name Name

func main() {
	fmt.Println("Please input your filename.")
	fmt.Scan(&fn)
	dat, _ := os.Open(fn)
	scanner := bufio.NewScanner(dat)
	name_slice := make([]Name, 0)
	for scanner.Scan() {
		string := string(scanner.Text())
		string_slice := strings.Fields(string)
		name.fname = string_slice[0]
		name.lname = string_slice[1]
		name_slice = append(name_slice, name)
	}
	fmt.Println(name_slice)
}
