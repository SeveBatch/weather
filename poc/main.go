package main

import (
	"fmt"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Good Stevening!")

}

func main() {
	//register generic handler
	http.HandleFunc("/", handler)
	log.Println("Server starting on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

