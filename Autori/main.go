package main

import (
	"fmt"
	"strings"
)

func main() {
	var names string

	//get names
	fmt.Scanln(&names)

	//split in strings
	substrings := strings.Split(names, "-")

	firstLetters := make([]string, len(substrings))

	for i, substring := range substrings {
		firstLetters[i] = strings.ToUpper(substring[0:1])
	}

	shortname := strings.Join(firstLetters, "")

	fmt.Println(shortname)
}
