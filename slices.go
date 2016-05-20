package main

import (
	"fmt"

)

func main() {

	//Declares a slice called myCourses
	// make(<type>, <len>, <cap>)
	// myCourses := make([]string, 5, 10)

	myCourses := []string{"Docker", "Puppet", "Python"}

	fmt.Printf("Length is: %d. \nCapcity is: %d",
		len(myCourses), cap(myCourses))

}
