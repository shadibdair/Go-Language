package main

import "fmt"

func main(){
	
	EvenNum := [5]int{0,2,4,6,8}
	for _,value := range EvenNum{
		fmt.Println(value)
	} 

	for i,value := range EvenNum{
		fmt.Println(value,i)
	} 

	numSlice := []int{5,4,3,2,1}
	sliced := numSlice[3:5]
	fmt.Println(sliced) // [2 1]

	slice2 := make([]int,5,10)
	copy(slice2,numSlice)
	fmt.Println(slice2) // [5 4 3 2 1]

	slice3 := append(numSlice, 3,0,-1)
	fmt.Println(slice3) // [5 4 3 2 1 3 0 -1]
}

// 0
// 2
// 4
// 6
// 8

// 0 0
// 2 1
// 4 2
// 6 3
// 8 4