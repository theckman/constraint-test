# dependency

This approach relies on the new package I've worked on, `goconstraint`, to show
a reusable method for defining build constraints without needing to make changes
to your code layout.

I'd suggest looking at the `simple` and `simple2` approaches to see how this is
an improvement.

Unlike the other two approaches, this mechanism uses a single blank import to
indicate that our software requires Go 1.9+. A good practice would be to declare
why we have the Go dependency:

```Go
// XXX(heckman): this depends on Go 1.9+, see this issue for more details:
//
// - <URL>
import _ "github.com/theckman/goconstraint/go1.9/gte"
```

When building this directory, as-is, the build failure looks like this:

```
# github.com/theckman/constraint-test/vendor/github.com/theckman/goconstraint/go1.9/gte
../vendor/github.com/theckman/goconstraint/go1.9/gte/constraint.go:10: undefined: __SOFTWARE_REQUIRES_GO_VERSION_1_9__
```

This approach is very similar to the `simple2` approach, except it doesn't
require that any functions in the file requiring Go 1.9 support be redefined to
avoid confusing compilation errors. The package also uses descriptive constants
so that developers/operators should be able to understand why the build clearly
failed.
