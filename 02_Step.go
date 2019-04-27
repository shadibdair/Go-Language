package main

import "fmt"

func main(){

	x,y := 12,6

	fmt.Println("x + y = ",x+y)
	fmt.Println("x - y = ",x-y)
	fmt.Println("x * y = ",x*y)
	fmt.Println("x / y = ",x/y)
	fmt.Println("x % y = ",x%y)


	isbool := true
	hate := false
	
	fmt.Println(isbool && hate)
	fmt.Println(isbool || hate)
	fmt.Println(!isbool)

}

/*

x + y =  18
x - y =  6
x * y =  72
x / y =  2
x % y =  0

false
true
false

*/