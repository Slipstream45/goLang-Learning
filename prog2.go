package main

import (
	"fmt"
)

func main() { 

	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("Now I am left with", originalCount-eatenCount, "apples.")
}

//can also do following

package main

import (
	"fmt"
)

func main() {

	originalCount := 10
	eatenCount := 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("Now I am left with", originalCount-eatenCount, "apples.")
}

