package main

import (
	"fmt"
)

func main() {

	name := "Ross"      // Name of subscriber
	course := "Go Lang" // Name of current course
	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {

	*course = "First Look: Native Docker Clustering"

	fmt.Println("Trying to change your course to", *course)

	return *course
}
