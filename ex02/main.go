package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Products struct {
	Id int
	Title string
	Price float32
	Description string
	Category string
	Image string
	Rating RatingProducts
}

type RatingProducts struct {
	Rate float32
	Count int
}

func printObject(object Products, arg string) {
	if arg == "id" {
		fmt.Println("id :", object.Id)
	} else if arg == "title" {
		fmt.Println("title: ",object.Title)
	} else if arg == "price" {
		fmt.Println("price: ", object.Price)
	} else if arg == "description" {
		fmt.Println("description: ", object.Description)
	} else if arg == "category" {
		fmt.Println("category: ", object.Category)
	} else if arg == "image" {
		fmt.Println("image: ", object.Image)
	} else if arg == "rating" {
		fmt.Println("rating: ", object.Rating.Rate)
	} else if arg == "count" {
		fmt.Println("count: ", object.Rating.Count)
	}
}

// if the arg is a valid one return true 

func findElement(arg string) bool {
	var types []string =  []string {
		"id", "title", "price", "description",
		"category", "image", "rate", "count"}
	for _, value := range types {
		if value == arg {
			return true
		}
	}
	return false
}

// check if is a number

func checkNumber(number string) int {
	intVar, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return intVar
}

// check if is a valid argument

func checkArgs(args []string) {
	for _, value := range args {
		if findElement(value) == false {
			fmt.Println("Invalid argument")
			os.Exit(1)
		}
	}
}

// get response data from url

func getRespose(url string, args []string) []byte {
	checkNumber(args[0])
	response, err := http.Get(url + args[0])
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	args = args[1:]
	checkArgs(args)
	responseData, err := ioutil.ReadAll(response.Body)
	
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return responseData
}

// Print all data
func printData(responseData []byte, args []string, responseObject []Products) {
	args = args[1:]
	err := json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, value := range responseObject {
		for _, arg := range args {
			printObject(value, arg)
		}
		fmt.Println("")
	}
}

func main() {
	var responseObject []Products
	args := os.Args[1:]
	responseData := getRespose("https://fakestoreapi.com/products?limit=", args)
	printData(responseData, args, responseObject)

}
