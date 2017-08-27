# simple

This path provides a simple approach to doing the version compatibility
constraint. Specifically, it guards the file that claims to have a Go 1.9
dependency with a build tag:

```Go
// +build go1.9
```

When building this directory, as-is, the build failure looks like this:

```
# github.com/theckman/constraint-test/simple
./simple.go:6: undefined: stringForThing
./simple.go:7: undefined: otherStringThing
```

When the build tag's condition isn't met, the source of that file is not
compiled. Unfortunately, this the build errors aren't helpful because they
indicate that an identifier is undefined, but the compiler doesn't explain why
they are undefined. In this case, they are undefined because our Go runtime
version is too old.

The `simple2` package builds on this idea to have the build failures be a bit
more helpful. Please see that package's `README.md` file for more info.
