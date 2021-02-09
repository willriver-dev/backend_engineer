package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleConnection)
	if error := http.ListenAndServe(":8080", nil); error != nil {
		log.Fatal(error)
	}
}

func handleConnection(write http.ResponseWriter, req *http.Request) {
	fmt.Printf("Http request header \n")
	for name, values := range req.Header {
		for _, value := range values {
			fmt.Printf("******************\n")
			fmt.Println(name, value)
		}
	}

	fmt.Printf("\n Http request method %s \n", req.Method)
	fmt.Printf("\n Http request body %s \n", req.Body)
	fmt.Printf("=================================\n", req.Body)

	switch req.Method {
	case "GET":
		fmt.Fprintf(write, "<h1> Hello World <h1>")

	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(write, "ParseForm error : %v", err)
		}

		fmt.Fprintf(write, "Post from website: %v\n", req.PostForm)
		name := req.FormValue("name")
		address := req.FormValue("address")
		fmt.Fprintf(write, "Name: %s, Address: %s\n", name, address)

	default:
		fmt.Fprintf(write, "Sorry only GET and POST method are supported")
	}
}
