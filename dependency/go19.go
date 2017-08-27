// This source code is released as-is in to the Public Domain.

package main

import (
	"fmt"

	// XXX(heckman): this depends on Go 1.9+, see this issue for more details:
	//
	// - <URL>
	_ "github.com/theckman/goconstraint/go1.9/gte"
)

func stringForThing() string {
	return "go1.9"
}

func otherStringThing(s string) string {
	return fmt.Sprintf("other string thing: %s", s)
}
