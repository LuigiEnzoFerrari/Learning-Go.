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

func checkNumber(age string) int {
	intVar, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return intVar
}

func checkArgs(args []string) {
	for _, value := range args {
		if findElement(value) == false {
			fmt.Println("Invalid argument")
			os.Exit(1)
		}
	}
	return
}

func main() {
	args := os.Args[1:]
	checkNumber(args[0])
	response, err := http.Get("https://fakestoreapi.com/products?limit=" + args[0])
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

	var responseObject []Products
	json.Unmarshal(responseData, &responseObject)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, value := range responseObject {
		for _, arg := range args {
			printObject(value, arg)
		}
	}
}
