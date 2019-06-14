package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var (
		slice  []int
		buffer bytes.Buffer
	)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Please, enter the number >:")
		scanner.Scan()
		buffer.Reset()

		buffer.Write(scanner.Bytes())

		switch i, err := strconv.Atoi(buffer.String()); {
		case err == nil:
			fmt.Printf("You entered the number [%d]\n", i)
			slice = append(slice, i)
		case buffer.String() == "X":
			fmt.Println("You entered the character ‘X’ for quit the program!\nSo, Bye-bye...")
			//Sorting and print slice
			printResult(slice)
			return
		default:
			fmt.Printf("Entered symbol [%s] is not int: %s\n", buffer.String(), err.Error())
		}
	}

}

func printResult(result []int) {
	sort.Ints(result)
	println("Printing entered result...")
	fmt.Printf("The slice has len=%d, cap=%d and include next elements: %v\n", len(result), cap(result), result)
}
