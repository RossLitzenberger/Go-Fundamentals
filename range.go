package main

import (
	"fmt"
)

func main() {

	coursesInProg := []string{"Docker Deep Dive", "Docker Clustering", "Docker and Kubernetes"}
	coursesCompleted := []string{"Docker Deep Dive", "Go Fundamentals", "Puppet Fundamentals"}

	// for _, i := range coursesInProg {
	// 	fmt.Println(i)
	// }
	for _, i := range coursesInProg {
		for _, j := range coursesCompleted {
			fmt.Println(j)
			if i == j {
				fmt.Println("\nHey we found a clash.", i , "is in both lists")
			}
		}
	}
}
