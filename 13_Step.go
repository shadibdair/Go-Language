package main

import "fmt"

func main(){
	rect1 := Rectangle{10,5}

	fmt.Println(rect1.height) // 10
	fmt.Println(rect1.width) // 5

	fmt.Println("Area of Rectangle is: ",rect1.area()) // Area of Rectangle is: 50
}

type Rectangle struct{
	height float64
	width float64
}

func (rect *Rectangle) area() float64{
	return rect.height * rect.width
}