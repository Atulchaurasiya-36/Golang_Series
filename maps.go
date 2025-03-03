package main

import "fmt"

func main(){
	// map declaration
	// using make() method
studentMarks:=make(map[string]int)
studentMarks["Atul"]=80
studentMarks["Vikrant"]=90
studentMarks["Rahul"]=70
fmt.Println(studentMarks)

// using map literal
moviesRating:=map[string]float64{
	"DDLJ":4.5,
	"3 idiots":4.8,
	"sholay":4.9,}
	fmt.Println(moviesRating)

	// accessing elements
	fmt.Println(studentMarks["Atul"])

	// updating elements
  studentMarks["Atul"]=90;
	fmt.Println(studentMarks)

	// deleting elements
	delete(moviesRating,"DDLJ")
	fmt.Println(moviesRating)

	// check if a key exists
	value ,exists:=moviesRating["3 idiots"]
	if exists{
		fmt.Println("movie rating exists",value)

	}else{
		fmt.Println("movie rating does not exists")
	}

	// iterating over the map

	for value ,movieName:=range moviesRating{
		fmt.Println(value,movieName)
	}

	



}