package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/alcor67/repo-Go-level-2-home-work/projectcalculator/calculator"
)

var op = flag.String("op", "", "operation + - * div % **")
var x = flag.Float64("x", 0, "first operand")
var y = flag.Float64("y", 0, "seconfd operand")
var result float64
var err error

func main() {
	flag.Parse()
	if *op == "" && *x == 0 && *y == 0 {

		fmt.Print("введите через пробел 2 значения и оператор\n",
			"{значение_1 оператор значение_2}\n")
		var numberString1, numberString2 string
		var operator string
		var number1, number2 float64

		fmt.Scanln(&numberString1, &operator, &numberString2)

		number1, err = strconv.ParseFloat(numberString1, 64)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		number2, err = strconv.ParseFloat(numberString2, 64)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		x = &number1
		y = &number2
		op = &operator
	}

	result, err = calculator.Calc(*op, *x, *y)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf(" результат: %v %s %v = %.6g \n", *x, *op, *y, result)
}
