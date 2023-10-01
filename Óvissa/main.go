package main

import (
	"fmt"
)

func main() {
	var uncert string
	var counter int
	fmt.Scanf("%s", &uncert)

	for _, letter := range uncert {
		if string(letter) == "u" {
			counter++
		}
	}
	fmt.Println(counter)
}
