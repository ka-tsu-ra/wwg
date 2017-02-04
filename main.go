package main

import (
	"encoding/json"
	"net/http"

	"github.com/katsuraku/wwg/animals"
)

// Something that's of interface is always a reference
var pets []animals.Pet

func main() {
	pets = []animals.Pet{
		&animals.Kitten{
			Name: "Mr Tiggles",
			Hobbies: []string{
				"Playing with wool",
				"Eating",
			},
		},
		&animals.Kitten{
			Name: "Mr Tom",
			Hobbies: []string{
				"Sleeping",
			},
		},
		&animals.Dog{
			Name: "Whatever",
			Hobbies: []string{
				"Barking",
				"Running",
			},
		},
	}
	http.HandleFunc("/list", ListPets)

	http.ListenAndServe(":9000", http.DefaultServeMux)
	// Could specify an IP instead of port number if you want
	// DefaultServeMux is a routing handler. Routes will be registered there.
}

func ListPets(rw http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(pets) // Need to convert Go object into Json to be rendered
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write(data)
}
