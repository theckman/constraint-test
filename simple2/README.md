# simple2

This path provides a second simple approach to doing the version compatibility
constraint. It builds on the `simple` approach, in that it tries to enhance the
build failure to include information to indicate why the build failed.

Like with `simple`, there is a file that will only be built with Go 1.9+

```Go
// +build go1.9
```

Unlike the first `simple` approach, this file also includes a defined constant
(`softwareRequiresGo19`) that we can reference elsewhere, allowing us to know
whether Go 1.9+ is being used.

The `simple2` approache also includes a `go19_nobuild.go` file which has the
following build tag to have it build when we're not using Go 1.9+:

```Go
// +build !go1.9
```

This `go19_nobuild.go` file also defines empty versions of the functions that
were undefined in the `simple` approach to avoid build failures that are
confusing:

```go
func stringForThing() string {
	panic("Go runtime version unsupported, Go 1.9+ required")
}

func otherStringThing(string) string {
	panic("Go runtime version unsupported, Go 1.9+ required")
}
```


When building this directory, as-is, the build failure looks like this, which is
an improvement over `simple`:

```
# github.com/theckman/constraint-test/simple2
./simple2.go:16: undefined: softwareRequiresGo19
```

When the Go 1.9+ build condition isn't met, the `softwareRequiresGo19` constant
does not get defined. This means any code referencing this constant at compile
time will fail. Because the constant name is descriptive, it's a more useful
error message for developers.

The biggest downside to the `simple2` approach is needing to redefine all of
your functions, in the guarded file, to avoid compile errors that are confusing.

The final example package, `dependency`, looks at using a blank import to
accomplish the same constraint checks.
