package main

import (
	"fmt"
	"net/http"
)

// can only have one func main in a Go program
// it is the entry point - the thing that will run

func main() {
	http.HandleFunc("/hello", HelloWorld)
	// creates /hello route. Whenever a request comes in on /hello, execute HelloWorld function
	http.HandleFunc("/goodbye", Goodbye)
	http.ListenAndServe(":9000", http.DefaultServeMux)
	// Could specify an IP instead of port number if you want
	// DefaultServeMux is a routing handler. Routes will be registered there.
}

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // get the value from the querystring that corresponds with 'name' (http://localhost:9000/hello?name=kirsten)
	fmt.Fprint(rw, "Hello "+name)
}

func Goodbye(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(rw, "Goodbye "+name)
}
