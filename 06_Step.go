package main

import "fmt"

func main(){
	age := 18

	if age > 18{
		fmt.Println("Yes, you can vote.")
	}else{
		fmt.Println("No, you can't vote.")
	}


	switch age{
	case 16: fmt.Println("Prepare for collage")
	case 18: fmt.Println("don't run after girls")
	case 20: fmt.Println("get your self a job!")
	default: fmt.Println("Are you even alive?")
	}
}

// No, you can't vote.

// don't run after girls