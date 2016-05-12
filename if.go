package main

import (
	"fmt"
)

func main() {
	// Variables to store course rankings
	firstRank := "39"   //Docker Deep Dive
	secondRank := "614" //Docker Clustering

	if firstRank < secondRank {
		fmt.Println("\n First course is doing better" + " then second course")
	} else if firstRank > secondRank {
		fmt.Println("\nOh dear... your first" + "course must be doing abysmally!")
	} else {
		fmt.Println("\nBoth courses are either" + "the same or something weird is going on")
	}
}
