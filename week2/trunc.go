package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Enter a floating point number :> ")
	var input string
	if _, err := fmt.Scanf("%s \n", &input); err == nil {
		if fl, err := strconv.ParseFloat(input, 64); err == nil {
			fmt.Printf("Printing truncated integer: %d \n", int64(fl))
		} else {
			fmt.Println("Wrong input! \\n Use float number for input...")
		}
	}
}
