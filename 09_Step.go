package main

import "fmt"

func main(){
	x,y := 5,3

	fmt.Println(add(x,y)) // 8
}

func add(num1 , num2 int)int{
	return num1+num2
}