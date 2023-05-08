package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// need to know diff between slice and array
// slice: pointer, size, capacity
var myblob = []int{1, 5, 21, 42, 69, -1, 0}

// read up on go escape analysis
// stack - comes with the function call, return val
func toJson(w http.ResponseWriter, _ *http.Request) {
	mynewblob := append(myblob, 15)
	mynewblob2 := append(mynewblob, 88, 127)
	mynewblob[0] = 2
	mynewblob = mynewblob[1:3]
	mynewblob = append(mynewblob, -100)
	json.NewEncoder(w).Encode(mynewblob)
	fmt.Printf("myblob %v\n", myblob)
	fmt.Printf("mynewblob %v\n", mynewblob)
	fmt.Printf("mynewblob2 %v\n", mynewblob2)
}

// Race condition - output depends on ordering of components
// data race - a special kind of race condition where a shared mutable variable may be updated in an (unexpected) order
func main() {
	http.HandleFunc("/", toJson)
	fmt.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil) // concurrent!
}
