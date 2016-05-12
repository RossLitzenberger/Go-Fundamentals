package main

import (
	"fmt"
)

func main() {

	if firstRank, secondRank := 39,
		614; firstRank < secondRank {
		fmt.Println("\n First course is doing better" +
			 	" then second course")
		if firstRank > 100 {
			fmt.Println("Error 1")
		}
	} else if firstRank > secondRank {
		fmt.Println("\nOh dear... your first" +
			 	"course must be doing abysmally!")
			if secondRank > 100 {
				fmt.Println("Error 2")
			}
	} else {
		fmt.Println("\nBoth courses are either" +
			 	"the same or something weird is going on")
	}
}
