package main

import (
	"biblioteki/engine"
	"bytes"
	"cmp"
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {
	calculator := engine.Calculator{}
	hashed := engine.Name{}
	var a, b int
	var name string
	fmt.Println("What's yout name?")
	fmt.Scanln(&name)
	fmt.Printf("Welcome to the Simple Calculator Program, %s!", name)
	fmt.Print("\nEnter the first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b)

	fmt.Println("*If: first is less then second you'll see red,\n" +
		"first equals second you'll see yellow\n" +
		"first is greater then second you'll see pink*")

	calculator.Add(a, b)
	calculator.Subtraction(a, b)
	calculator.Multiplication(a, b)
	calculator.Division(float64(a), float64(b))
	calculator.Poww(float64(a), float64(b))
	hashed.Hashed(name)
	buf := bytes.Buffer{}
	buf.WriteString(time.Now().Local().String())
	switch {
	case bytes.Contains(buf.Bytes(), []byte("2024")):
		color.Green("\nIt's the year 2024, exact time is: %s, what would you like to do?", buf.String())
		color.Blue("There are some operations on numbers %d and %d", a, b)
		switch {
		case cmp.Compare(a, b) == -1:
			color.Red("Sum: %d, Sub: %d, Mult: %d, Div: %.2f, Pow: %.2f", calculator.Sum, calculator.Sub, calculator.Mult, calculator.Div, calculator.Pow)
		case cmp.Compare(a, b) == 0:
			color.Yellow("Sum: %d, Sub: %d, Mult: %d, Div: %.2f, Pow: %.2f", calculator.Sum, calculator.Sub, calculator.Mult, calculator.Div, calculator.Pow)
		case cmp.Compare(a, b) == 1:
			color.Magenta("Sum: %d, Sub: %d, Mult: %d, Div: %.2f, Pow: %.2f", calculator.Sum, calculator.Sub, calculator.Mult, calculator.Div, calculator.Pow)
		}

	default:
		panic("You've distorted reality!")
	}
	fmt.Printf("And as a curiosity there is your name in md5 format: %s", hashed.Hash)

}
