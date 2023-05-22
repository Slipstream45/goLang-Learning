package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the game!")
	target := int64(rand.Intn(100))
	status := 0

	for i := 0; i < 10; i++ {
		fmt.Println("Guess #", i)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = strings.TrimRight(input, "\r\n")
		guess, _ := strconv.ParseInt(input, 10, 32)
		if guess > target {
			fmt.Println("OOPS, your guess was HIGH!")
		} else if guess < target {
			fmt.Println("OOPS your guess was LOW")
		} else {
			fmt.Println("CORRECT GUESS!!")
			status = 1
			break
		}

	}
	if status == 0 {
		fmt.Println("Sorry you didn't guess my number! It was:", target)
	}
}
