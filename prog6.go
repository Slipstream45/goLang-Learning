package main

import (
	"fmt"
	"math"
)

func main() {
	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)
}

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}


//another one

package main

import (
	"fmt"
)

func main() {
	amount, err := paintNeeded(4.2, -3.0)
	if err == nil {
		fmt.Printf("%0.2f litres required!", amount)
	} else {
		fmt.Println(err)
	}
}
func paintNeeded(length float64, breadth float64) (float64, error) {
	if length < 0 || breadth < 0 {
		return 0, fmt.Errorf("Invalid parameters given!!")
	}
	area := length * breadth
	return area, nil
}
