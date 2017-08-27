// This source code is released as-is in to the Public Domain.

package main

import "fmt"

func worksWithAll() string {
	thing := stringForThing()
	return otherStringThing(thing)
}

func main() {
	fmt.Println(worksWithAll())
}
