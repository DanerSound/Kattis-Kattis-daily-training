package main

import "fmt"

func main(){
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2, &num3)
	if (num1 + num2 == num3){
		fmt.Printf("correct!")
	}else{
		fmt.Printf("wrong!")
	}
}