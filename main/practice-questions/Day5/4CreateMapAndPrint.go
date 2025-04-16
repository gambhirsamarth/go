package main

import "fmt"

func main() {
	cityToPopulationMap := map[string]int{
		"Delhi":      21389682473,
		"Chandigarh": 894297964,
		"Rishikesh":  42759653,
	}

	//for city := range cityToPopulationMap {
	//	fmt.Println(city, cityToPopulationMap[city])
	//}

	for city, population := range cityToPopulationMap {
		fmt.Printf("City: %s, Population: %d\n", city, population)
	}

}
