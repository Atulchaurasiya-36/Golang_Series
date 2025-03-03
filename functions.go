package main

import "fmt"

// function without argument & return 

func greet(){
	fmt.Println("hello")
}

// function with argument 
func sayHello(greet string){
	fmt.Println(greet)
}

// function with argument and return 

// func add(a,b int)int{
// 	return a+b;
// }

// function with multiple return values

func divide(a , b int)(int ,int){
	quotient:=a/b
	remainder:=a%b;
	return quotient, remainder
}

// a function take argument to another function

func applyOperation(a,b int,operation func(int ,int)int)int{
	return operation(a,b)
}

func add(a,b int)int{
	return a+b;
}

func main(){
	// greet();
	sayHello("hello Atul")
	result:=add(5,90)
	fmt.Println(result)
	q,r :=divide(10,20)
	fmt.Println(q,r)
	result1:=applyOperation(12,20,add)
	fmt.Println(result1)
	

}