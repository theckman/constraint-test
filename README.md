# constraint-test

A repo with some different ways of enforcing Go constraints to show tradeoffs.
Each directory has a `README.md` file in it to provide more detail on that
approach. Let's do a quick overview:

## License

All of the source code, examples, and documentation in this repository are
released in to the Public Domain using the Unlicense. You are free to use,
modify, distribute, and break this code.

## Approaches

Let's take a look at a few different ways to solve this problem, and why it may
not be ideal.

### simple

This is the most simple approach of guarding your file with build tags (`//
+build go1.9`). It shows that the compilation errors can be confusing, if you
decide to restrict to a specific Go runtime, since the reason for the build
failure isn't clearly indicated. Instead of saying that we need Go 1.9+, it is
showing us that identifiers are undefined:

```
# github.com/theckman/constraint-test/simple
./simple.go:6: undefined: stringForThing
./simple.go:7: undefined: otherStringThing
```

### simple2

This is very similar to `simple`, except we also add a new `const` to indicate
Go 1.9 support and we duplicate all of the missing functions in to a file
guarded by the inverse build tag (`// +build !go1.9`). While this mechanism
improves the build failure message, it requires that any functions used by other
files be duplicated in to the dummy file. The resulting build error looks like
this:

```
# github.com/theckman/constraint-test/simple2
./simple2.go:16: undefined: softwareRequiresGo19
```

### dependency

The third approach is one I'm calling `dependency`, as it introduces a new
package for doing build constraints using blank imports:

* https://github.com/theckman/goconstraint

It uses blank imports at the top of your source file, which you can annotate
with comments to provide context:

```Go
// XXX(heckman): this depends on Go 1.9+, see this issue for more details:
//
// - <URL>
import _ "github.com/theckman/goconstraint/go1.9/gte"
```

When building the code against a version of Go that isn't supported, the build
failure is also descriptive:

```
# github.com/theckman/constraint-test/vendor/github.com/theckman/goconstraint/go1.9/gte
../vendor/github.com/theckman/goconstraint/go1.9/gte/constraint.go:10: undefined: __SOFTWARE_REQUIRES_GO_VERSION_1_9__
```
