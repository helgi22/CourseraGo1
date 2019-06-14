package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var slice []int

	for {
		var input string

		fmt.Printf("Enter yours number:")
		_, _ = fmt.Scanln(&input)

		switch num, err := strconv.Atoi(input); {
		case strings.ToUpper(input) == "X":
			fmt.Printf("You entered the character ‘X’ for quit the program!\n"+
				"Final result: Slise has len=%d, cap=%d and it include next elements: %v\n"+
				"So, Bye-bye...", len(slice), cap(slice), slice)
			return
		case err == nil:
			slice = append(slice, num)
			sort.Ints(slice)
			fmt.Printf("Slice include: %v\n", slice)
		default:
			fmt.Printf("Entered symbol [%s] is not int number: %s\n", input, err.Error())
		}

	}
}
