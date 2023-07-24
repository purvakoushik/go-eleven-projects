package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static-html-images"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server is going to listen at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 Not Found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "405 Method not allowed ", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Hello, How are you my friend ?")
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(res, "Post form succeeded \n")
	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintf(res, "name : %v \n", name)
	fmt.Fprintf(res, "address : %v \n", address)
}
