package main

import "fmt"

func main(){
	FirstRun()
	SecondRun()
}

func FirstRun(){fmt.Println("I excute First")}
func SecondRun(){fmt.Println("I excute Second")}


// I excute First
// I excute Second

// with defer
// I excute Second
// I excute First
