package main

import (
	"fmt"
	"os"
)

func main() {

	name := os.Getenv("USER") // Name of User
	course := "Go Lang"       // Name of current course

	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {

	*course = "First Look: Native Docker Clustering"

	fmt.Println("Trying to change your course to", *course)

	return *course
}
