package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guimassoqueto/go_web_app/pkg/handlers"
)

const portNumber string = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server is listening on port %s...", portNumber)

	err := http.ListenAndServe(portNumber, nil)

	if err != nil {
		log.Fatal(err)
	}
}
