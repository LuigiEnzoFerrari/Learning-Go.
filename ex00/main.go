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

//create the number of persons as the number passed in the first arg
func createPerson(args []string, number int) []Person {
	checkArgs(args, number)
	var persons []Person
	for i := 0; i < number; i++ {
		persons = append(persons,
			Person{args[i], checkNumber(args[i + number])})
	}
	fmt.Println(persons)
	return persons
}

// check if the args passed is a number
func checkNumber(age string) int {
	intVar, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return intVar
}

// check the number of arguments
func checkArgs(args []string, number int) {
	fmt.Println(len(args))
	if len(args) < number * 2 {
		fmt.Println("Error: Not enough arguments")
		os.Exit(1)
		
		} else if len(args) > number * 2 {
			fmt.Println("Error: Too many arguments")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Sintex: <number of members> <name> <age>")
	args := os.Args[1:]
	createPerson(args[1:], checkNumber(args[0]))
}
