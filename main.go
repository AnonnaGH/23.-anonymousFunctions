package main

import "fmt"


func main() {
	a := func(x,y int)(r int){
	r=x+y
	return
	}(10,10)
	
	fmt.Println(a)
}