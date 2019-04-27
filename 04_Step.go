package main

import "fmt"

func main(){

	var name string = "Shadi Bdair"
	const pi float64 = 4.12312432
	win := true
	x := 5

	fmt.Println(len(name)) // 11
	fmt.Println(name + " the best one!!") // Shadi Bdair he best one!!

	fmt.Printf("%.3f \n",pi) // 4.123
	fmt.Printf("%T \n",name) // string
	fmt.Printf("%t \n",win) // ture
	fmt.Printf("%d \n",x) // 5
	fmt.Printf("%b \n", 25) // 11001
	fmt.Printf("%c \n", 33) // "!"
	fmt.Printf("%x \n", 12) // "c"
	fmt.Printf("%e \n", pi) // 4.123124e+00
}