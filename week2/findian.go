package main

import (
	"fmt"
	"regexp"
)

func main() {
	var validStr = regexp.MustCompile(`^(I|i){1}([\S\s]*[aA]+[\S\s]*)+(N|n){1}$`)
	fmt.Print("Please, enter your text:>")
	var input string
	if _, err := fmt.Scanf("%q", &input); err == nil {
		if validStr.MatchString(input) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Ops, something went wrong! Try to enclose in double quotes...", err)
	}
}
