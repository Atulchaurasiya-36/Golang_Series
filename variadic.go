package main

import "fmt"

// variadic function

func sum(nums...int)int{
	total:=0;
	for _,num:=range nums{
		total+=num;
	}
	return total
}

func main(){
	fmt.Println(sum(1,3,5,7,45,23,12,90))
}