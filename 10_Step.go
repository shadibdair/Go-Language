package main

import "fmt"

func main(){

	num := 5
	fmt.Println(factorial(num)) // 120 
}	

func factorial(num int) int{
	if num == 0{
		return 1
	}
	return num*factorial(num-1) 
}