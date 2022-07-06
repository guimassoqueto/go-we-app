package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber string = ":8000"

func Home(res http.ResponseWriter, req *http.Request) {
	size_response, err := fmt.Fprint(res, `<h1>Home Page</h1><a href="/about">Go to About Page</a>`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(size_response)
}

func About(res http.ResponseWriter, req *http.Request) {
	size_response, err := fmt.Fprint(res, `<h1>About Page</h1><a href="/">Go to Home</a>`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(size_response)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server is listening on port %s...", portNumber)

	err := http.ListenAndServe(portNumber, nil)

	if err != nil {
		log.Fatal(err)
	}
}
