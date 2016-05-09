package main

import (
	"fmt"
	"reflect"
)

var  (
	name = "Ross" // Name of subscriber
	course = "Go Lang" // Name of current course
	module = 3.2 // Current place in course
)

func main() {
	fmt.Println("Name is", name, "add is  of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "add is  of type", reflect.TypeOf(module))
}
