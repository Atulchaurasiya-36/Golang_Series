package main

import "fmt"
 func main(){
	// slice declaration and initialization
	number:=[]int{3,4,5,6,7}
	fmt.Println(number)

	// appending values in slice
	number=append(number,8,9,12,89,90)
	fmt.Println(number)

	// traversing on slice

	for i:=0;i<len(number);i++{
		fmt.Println(number[i])
	}

	// declaring slice using make()
	slice1:=make([]int ,2,5)
	slice1=append(slice1,1,2,3,4)
	// display len of slice
	fmt.Println(len(slice1))

	// display capacity of slice1
	fmt.Println(cap(slice1))

	// there is one more way for creating slice ->shortHand syntax

	nums:=[]int{}
	nums = append(nums, 1) 
	nums = append(nums, 4) 
	fmt.Println(len(nums))
	fmt.Println(cap(nums))



 }