package main

import "fmt"

// generic function

func gen[T any](items[]T){
	for _,item:=range items{
		fmt.Println(item);
	}
}

func main(){
	nums := []string{"golang", "js", "java", "python"}
	gen(nums)

	numbers:=[] int{2,4,5,6,7,8,9}
	gen(numbers)

}