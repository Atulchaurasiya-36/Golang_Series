package main

import "fmt"

// Function to receive data
func receiveData(ch chan string) {
	fmt.Println("Received:", <-ch) 
}

func main() {
	ch := make(chan string)
	go receiveData(ch) 
	ch <- "Hello Atul" 
}
