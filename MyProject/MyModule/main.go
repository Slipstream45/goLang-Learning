package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16 // min 0 max 65535
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func main() {
	a_car := car{22341, 0, 12561, 225.0}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
}
