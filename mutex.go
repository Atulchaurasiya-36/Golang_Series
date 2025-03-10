package main

import (
	"fmt"
	"sync"
	"time"
)

var counter=0
// declaring mutex
var mu sync.Mutex 
func incerement(){
	for i:=0;i<1000;i++{
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func main(){
	go incerement()
	go incerement()
	time.Sleep(time.Second)
	fmt.Println(counter)
}