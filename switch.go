package main

import (
	"fmt"
	// "time"
)

func main(){
	// switch case
	// i:=3
	// switch i{
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// default:
	// 	fmt.Println("others")
	// }

	// multiple conditions
	// switch time.Now().Weekday(){
	// case time.Saturday,time.Sunday:
	// 	fmt.Println("it's weekend")
	// default:
	// 	fmt.Println("it is workday")
		
	// }

whoAmI:=func(i interface{}){
	switch i.(type){
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	}
}
whoAmI(21)


}