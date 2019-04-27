package main

import "fmt"

func main(){
	
	StudentAge := make(map[string] int)

	StudentAge["Shadi"] = 21
	StudentAge["ronaldo"] = 31

	fmt.Println(StudentAge) // map[Shadi:21 ronaldo:31]
	fmt.Println(StudentAge["Shadi"]) // 21
	fmt.Println(len(StudentAge)) // 2


	superhero := map[string]map[string]string{
		"superman" : map[string]string{
			"RealName" :"Clark knet",
			"City" : "Metropolis",
		},

		"batman" : map[string]string{
			"RealName" :"Bruse Wayne",
			"City" : "Gotham City",
		},
	}

	if temp,hero := superhero["superman"]; hero{
		fmt.Println(temp["RealName"], temp["City"]) // Clark knet Metropolis
	}
}