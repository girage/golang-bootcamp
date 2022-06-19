package main

import "fmt"

func main() {
	speed := 100 // speed is int
	force := 2.5 //force is float64

	// speed = speed * force // error becuz diff. type
	// speed = speed * int(force) // => 200 (100 * 2)
	// speed = float64(speed) * force // error becuz you speed is int
	speed = int(float64(speed) * force) //  250

	fmt.Println(speed)
}
