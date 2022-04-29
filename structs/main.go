package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func checkAge(age string) int {
	intVar, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return intVar
}

func checkArgs(args []string, number int) {
	if len(args) - 2 < number * 2 {
		fmt.Println("Error: Not enough arguments")
		os.Exit(1)
		
		} else if len(args) - 2 > number * 2 {
			fmt.Println("Error: Too many arguments")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Sintex: <number of members> <name> <age>")
	myargs := os.Args[1:]
	var persons []Person
	number := checkAge(myargs[0])
	checkArgs(myargs, number)

	for i := 0; i < number; i++ {
		persons = append(persons,
				Person{myargs[i + 1], checkAge(myargs[i + number + 1])})
		fmt.Println(myargs[i+ 1])
	}
	fmt.Println(persons)
}
