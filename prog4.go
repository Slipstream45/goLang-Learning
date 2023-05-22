package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	status := ""
	fmt.Println("Enter a number")

	//this is how we take input
	reader := bufio.NewReader(os.Stdin)
	input, error1 := reader.ReadString('\n')

	//handle error from ReadString()
	if error1 != nil {
		log.Fatal(error1)
	}
	//How we trim and make the input which is a string to a float
	input = strings.TrimRight(input, "\r\n") // Assign the trimmed input back to the variable
	grade, error2 := strconv.ParseFloat(input, 64)
	//handle error here
	fmt.Println(grade)
	if error2 != nil {
		log.Fatal(error2)
	}

	//conditional to check
	if grade > 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("Grade is", grade, "\nStatus is :", status)
}
