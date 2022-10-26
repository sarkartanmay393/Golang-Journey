package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// By default, we choose 8080.
var portNumber = "8080"

// main creates a very simple web server that serve html from static directory.
func main() {
	// heroku provides open port by environment variable.
	portNumber, found := os.LookupEnv("PORT")
	if !found {
		log.Fatalf("Port env not found !")
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(fmt.Sprintf(":%v", portNumber), nil)
	if err != nil {
		log.Fatal(err)
	}
}
