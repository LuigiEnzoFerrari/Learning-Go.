package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
	"github.com/gorilla/mux"
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

var friends []Friends

// get all the friends data

func names(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("404 - not found\n"))
		return
	}

	var names []string
	for _, friend := range friends {
		names = append(names, friend.Name)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(names)
	log.Println(r.Method, r.URL.Path)
}

// get and post requests to /friends

func handleFriends(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/friends" {	
		w.WriteHeader(404)
		w.Write([]byte("404 - not found\n"))
		return
	}
	log.Println(r.Method, r.URL.Path)
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(friends)
	case "POST":
		var newfriends Friends
		json.NewDecoder(r.Body).Decode(&newfriends)
		friends = append(friends, newfriends)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(friends)
	}
}

// get and update the value on the server

func handleFriend(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	switch r.Method {
	case "PUT":
			var newfriend Friends
			json.NewDecoder(r.Body).Decode(&newfriend)
			for i, friend := range friends {
				if friend.Name == mux.Vars(r)["name"] {
					friends[i] = newfriend
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(friends[i])
					return
				}
			}
	case "GET":
		log.Println("we are in GET")
		for i, friend := range friends {
			if friend.Name == mux.Vars(r)["name"] {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(friends[i])
					return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - not found\n"))
	}
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Create initial data values
	friends = append(friends, Friends{Name: "John"})
	friends = append(friends, Friends{Name: "Luigi"})

	// Attach handlers to the router
	r.HandleFunc("/", names).Methods("GET")
	r.HandleFunc("/friends", handleFriends).Methods("GET", "POST")
	r.HandleFunc("/friends/{name}", handleFriend).Methods("GET", "PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
