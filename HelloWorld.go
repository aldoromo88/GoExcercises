package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var age = 40

	var favNum = 1.61355

	randNum := 1

	fmt.Println(age, favNum, randNum)

	var numOne = 1.000
	var num99 = 0.9999

	fmt.Println(numOne - num99)

	const pi float64 = 3.14159265

	var (
		varA = 2
		varB = 3
	)

	var result = varA + varB
	fmt.Println(result)

	var myName = "Aldo"
	fmt.Println(len(myName))

	var isOver40 = true

	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", pi)

	fmt.Printf("%d \n", 100)
	fmt.Printf("%b \n", 100)
	fmt.Printf("%c \n", 100)
	fmt.Printf("%x \n", 100)
	fmt.Printf("%e \n", pi)

	i := 1

	for i <= 10 {

		fmt.Println(i)
		i++

	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	if isOver40 {
		fmt.Println("Es mayor de 40")
	}

}
