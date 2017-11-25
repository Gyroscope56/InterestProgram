package main

import (
	"fmt"
)

func main() {
	x := 12
	Divisions := 1
	StartingValue := 10.0
	Percent := float64(100.0 / Divisions)
	Interest := StartingValue / 100 * Percent
	fmt.Println("Starting Value: $", StartingValue)
	fmt.Println("Interest: $", Interest)
	Total := (Interest + StartingValue)
	fmt.Println("Total Money after", x/Divisions, "months: $", Total)
	y := 2
	for i := 1; i < Divisions; i++ {
		Interest = Total / 100 * Percent
		fmt.Println("Interest after", x/Divisions*y, "months: $", Interest)
		Total = (Total + Interest)
		fmt.Println("Total after", x/Divisions*y, "months", ": $", Total)
		y++
	}
}
