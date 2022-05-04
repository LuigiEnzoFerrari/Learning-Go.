package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
 )
 
 type Friends struct {
	Name		string
	Age			int
	Body		Anatomy
	Anniversary	BirthDate
}

type BirthDate struct {
	Day		int
	Month	int
	Year	int	
}

type Anatomy struct {
	Height	float32
	Weight	float32
}

func getInput[T any](text string, friendVar *T) {
	fmt.Println(text)
	fmt.Scanln(friendVar)
} 

// get input from the user to create a new friend

func getUserInput() Friends {
	var userInput Friends
	getInput("Enter Friend's Name: ", &userInput.Name)
	getInput("Enter Friend's Age: ", &userInput.Age)
	getInput("Enter Friend's Height: ", &userInput.Body.Height)
	getInput("Enter Friend's Weight: ", &userInput.Body.Weight)
	fmt.Println("Enter Friend's Birth Date:")
	getInput("Enter Friend's Day: ", &userInput.Anniversary.Day)
	getInput("Enter Friend's Month: ", &userInput.Anniversary.Month)
	getInput("Enter Friend's Year: ", &userInput.Anniversary.Year)
	return userInput
}

func main() {
	// Encode the data
	userInput := getUserInput()
	postBody, _ := json.Marshal(userInput)
	// Create a request
	responseBody := bytes.NewBuffer(postBody)
	var method string
	fmt.Println("Enter Method: ")
	fmt.Scanln(&method)

	// send the request
	if (method == "POST") {
		resp, err := http.Post("http://localhost:8000/friends", "application/json", responseBody)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
	} else if (method == "PUT") {
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/friends/" + userInput.Name, responseBody)
		if err != nil {
			// handle error
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		resqs, err := client.Do(req)
		if err != nil {
			// handle error
			log.Fatal(err)
		}
		fmt.Println(resqs.Status)
	}
 }
