package main

import "fmt"

func main(){
	// Array declaration
	// method 1
	var arr0[5]int;
	arr0[0]=1;
	fmt.Println(arr0[0])

	// declaration and initialization
	var arr=[5]int{5,7,9,2,6};
	// traversal on array
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}

	// method 3
	arr1:=[5]int{9,0,4,3,1};
	fmt.Println(arr1);

	// declaring 2D array
	var twoD[2][3]int;
	twoD[0][0]=1;
	fmt.Println(twoD[0][0])
	// declaring and initializing 2D
	var twoD1=[2][3]int{{1,2,3},{4,7,8}}
	fmt.Println(twoD1)
	// traversing on 2D arrays

	for i:=0;i<len(twoD1);i++{
		for j:=0;j<len(twoD1[0]);j++{
			fmt.Println(twoD1[i][j])
		}
	}


}