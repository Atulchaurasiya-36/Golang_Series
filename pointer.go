package main

import "fmt"

func update(num *int){
	*num=10;
}

func main(){
	// declaration of pointer
   
	var ptr *int
	y:=90;
	ptr=&y;
	fmt.Println(*ptr)

// to pass address of variable to another function's variable

	x:=50;
	update(&x);
	fmt.Println(x)

}