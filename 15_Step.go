package main

import ("fmt"
		"os"
		"log"
         "io/ioutil")

func main(){
	file, err := os.Create("sample.txt")

	if err != nil{ // none
		log.Fatal(err)
	}

	file.WriteString("Hi, my name is Area and this file was created using GO!!!")
	file.Close()

	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil{ // none
		log.Fatal(err)
	}

	s1 := string(stream)
	fmt.Println(s1) // Hi, my name is Area and this file was created using GO!!!
}