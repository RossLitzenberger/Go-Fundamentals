package main

import (
	"fmt"
)

func main() {

	type courseMeta struct {
		Author string
		Level string
		Rating float64
	}

	// var DockerDeepDive courseMeta
	// DockerDeepDive := new(courseMeta)
	DockerDeepDive := courseMeta{
		Author: "Testing",
		Level: "Intermediate",
		Rating: 1,
	}
	fmt.Println("\nDocker Deep Dive author is:", DockerDeepDive.Author)
	fmt.Println("\nDocker Deep Dive rating is:", DockerDeepDive.Rating)
}
