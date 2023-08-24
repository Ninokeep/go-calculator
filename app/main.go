package main

import (
	"errors"
	"fmt"
)

func main(){

	var num1, num2 float64
	var operator string


	fmt.Print("Enter a number ...")
	fmt.Scan(&num1)

	fmt.Scan(&operator)

	switch operator {

	case "-":
		fmt.Scan(&num2)
		num1 -= num2
		fmt.Print(num1)

	case "*":
		fmt.Scan(&num2)
		num1 *= num2
		fmt.Print(num1)

	case "+":
		fmt.Scan(&num2)
		num1 += num2
		fmt.Print(num1)

	case "/":
		fmt.Scan(&num2)

		if(num2 == 0){
			fmt.Print(errors.New("You cannot divided by 0"))
			break
		}
		num1 += num2
		fmt.Print(num1)
	}



}



