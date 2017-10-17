package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Expecting just 1 argument")
		os.Exit(1)
	}
	fmt.Println("What it is first argument", os.Args[0])
	fmt.Println("It's over", os.Args[1])
}
