package main

import "fmt"

func main() {
	cityToPopulationMap := map[string]int{
		"Delhi":      21389682473,
		"Chandigarh": 894297964,
		"Rishikesh":  42759653,
	}

	cities := []string{"Delhi", "Chandigarh", "Rishikesh", "Bangalore", "Mumbai"}

	for city := range cities {
		population, ok := cityToPopulationMap[cities[city]]
		if ok {
			fmt.Printf("Population of %s is %d\n", cities[city], population)
		} else {
			fmt.Printf("Population of %s is not available\n", cities[city])
		}
	}
}
