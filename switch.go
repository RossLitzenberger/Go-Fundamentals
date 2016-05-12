package main

import (
	"fmt"
)

func main() {

	switch "docker" {
	case "linux":
		fmt.Println("\nHere are some recommanded" + " Linux courses...")
	case "docker":
		fmt.Println("\nHere are some recommanded" + " Docker courses...")
		fallthrough
	case "mac":
		fmt.Println("\nHere are some recommanded" + " Mac courses...")
	default:
		fmt.Println("\nSorry, we couldn't find a match," + "why not try our Top 100 list!")
	}
}
