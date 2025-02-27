package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("hello Atul");
	// taking input from user
	var name string
	fmt.Println("enter your name")
	fmt.Scan(&name);
	fmt.Println("your name is",name)

	// taking input from buffered reader
	reader:=bufio.NewReader(os.Stdin)
	name,_=reader.ReadString('\n')
	fmt.Printf("hello,%s",name)

}