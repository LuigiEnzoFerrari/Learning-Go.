package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
	"encoding/json"
	"io/ioutil"
)

type Person struct {
	Name	string
	Type	string
	Size	uint
	Age		uint
}

//set of names

var names [15]string =  [15]string {
	"Amen", "Brota", "Collen",
	"Detch", "Euler", "Fe",
	"Gustavo", "Henrique", "Igor",
	"Jose", "Kenny", "Lara",
	"Mateus", "Nathan", "Otavio"}

//set of names

var types [15]string =  [15]string {
	"Chad", "Alpha", "Gama",
	"Beta", "Omega", "Delta",
	"Nice", "Bad", "Neutral",
	"Chaotic", "Good", "Evil",
	"Funny", "Sexy", "Hot"}
	
// Random number generator
var myrand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomStrings(length int, sett [15]string) string {
	return sett[myrand.Intn(len(sett))]
}


// check the arguments
func checkArgs() int {
	Args := os.Args[1:]
	if len(Args) < 1 {
		fmt.Println("Error: Not enough arguments")
		os.Exit(1)
	}
	if len(Args) > 2 {
		fmt.Println("Error: Too many arguments")
		os.Exit(1)
	}
	number, err := strconv.Atoi(Args[0])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return number
}

// create a array of person with random values
func createPerson(number int) []Person {
	var persons []Person
	for i := 0; i < number; i++ {
		persons = append(persons,
			Person{RandomStrings(5, names),
				RandomStrings(5, types),
				uint(myrand.Intn(100)),
				uint(myrand.Intn(100))})
	}
	return persons
}

// create  a json file with the persons
func createJSon(persons []Person, filename string) {
	e, err := json.MarshalIndent(persons, "", "  ")
	_ = ioutil.WriteFile(filename, e, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	persons := createPerson(checkArgs())
	createJSon(persons, "persons.json")
}
