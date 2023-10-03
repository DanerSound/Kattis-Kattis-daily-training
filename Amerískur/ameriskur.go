package main

import "fmt"

func main() {
	var lengthInFields int
	fmt.Scan(&lengthInFields)
	lengthInKilometers := float64(lengthInFields) * 0.09144
	fmt.Printf("%f\n", lengthInKilometers)
}
