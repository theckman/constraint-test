// This source code is released as-is in to the Public Domain.
//
// +build !go1.9

package main

func stringForThing() string {
	panic("Go runtime version unsupported, Go 1.9+ required")
}

func otherStringThing(string) string {
	panic("Go runtime version unsupported, Go 1.9+ required")
}
