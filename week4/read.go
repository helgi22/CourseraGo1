package main

import (
	"bufio"
	"fmt"
	. "log"
	"os"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		Fatal(e)
	}
}

//
// Person struct which contains  the first name, and the last name
type Person struct {
	fName string
	lName string
}

func main() {
	var fileName string
	fmt.Printf("Enter full file name:")

	if _, err := fmt.Scan(&fileName); err != nil {
		fmt.Println("Wrong input!")
	}

	//Open file
	file, err := os.Open(fileName)
	checkErr(err)
	defer file.Close()

	fmt.Println("Successfully opened ", fileName)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var persons []Person
	for scanner.Scan() {
		splitWord := strings.Fields(scanner.Text())
		fName := splitWord[0]
		lName := splitWord[1]
		//restrict length of fName
		if len(fName) > 20 {
			fName = fName[:20]
		}
		//restrict length of lName
		if len(lName) > 20 {
			lName = lName[:20]
		}
		person := Person{fName: fName, lName: lName}
		fmt.Printf("Structure Person added to slice %v\n", person)
		persons = append(persons, person)
	}

	fmt.Println("\nPrint slice with all structures Person:")
	for key, value := range persons {
		fmt.Println(key, value)
	}
}
