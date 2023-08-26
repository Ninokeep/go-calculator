package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main(){

	var num1, num2 float64
	var operator string
	fmt.Println("Enter a number ...")

	fmt.Scan(&num1);

	for {

	fmt.Scanln(&operator);

	if operator == " " {
		break;
	}

	switch operator {

	case "-":
		fmt.Scanln(&num2);
		num1 -= num2;
		getTheResult(num1)

	case "*":
		fmt.Scanln(&num2);
		num1 *= num2;
		getTheResult(num1)
	case "+":
		fmt.Scanln(&num2);
		num1 += num2;
		getTheResult(num1)

	case "/":
		fmt.Scanln(&num2);

		if(num2 == 0){
			fmt.Print(errors.New("You cannot divided by 0"));
			break;
		}
		num1 /= num2;
		getTheResult(num1)
	}
	}

}

func getTheResult(num float64)  {

 fmt.Printf("Total : %s", strconv.FormatFloat(num, 'f', -1, 64));
}
