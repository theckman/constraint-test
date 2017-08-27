// This source code is released as-is in to the Public Domain.
//
// +build go1.9

package main

import "fmt"

const softwareRequiresGo19 = uint8(0)

func stringForThing() string {
	return "go1.9"
}

func otherStringThing(s string) string {
	return fmt.Sprintf("other string thing: %s", s)
}
