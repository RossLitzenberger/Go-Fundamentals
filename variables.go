package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Ross"      // Name of subscriber
	// course := "Go Lang" // Name of current course
	module := 3.2       // Current place in course
	ptr := &module

	fmt.Println("Name is", name, "add is  of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "add is  of type", reflect.TypeOf(module))
	fmt.Println("Memory address of *module* variable is ", &module)
	fmt.Println("Memory address of *module* variable is ", ptr, "add the value of *module* is", *ptr)

}
