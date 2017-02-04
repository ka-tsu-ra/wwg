package main

import (
	"encoding/json"
	"net/http"

	"github.com/katsuraku/wwg/animals"
)

var kittens []animals.Kitten

func main() {
	kittens = []animals.Kitten{
		animals.Kitten{
			Name: "Mr Tiggles",
			Hobbies: []string{
				"Playing with wool",
				"Eating",
			},
		},
		animals.Kitten{
			Name: "Mr Tom",
			Hobbies: []string{
				"Sleeping",
			},
		},
	}
	http.HandleFunc("/list", ListKittens)

	http.ListenAndServe(":9000", http.DefaultServeMux)
	// Could specify an IP instead of port number if you want
	// DefaultServeMux is a routing handler. Routes will be registered there.
}

func ListKittens(rw http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(kittens) // Need to convert Go object into Json to be rendered
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Write(data)
}
