Assert for Go testing
=====================

PkgDoc: [http://godoc.org/github.com/chai2010/assert.go](http://godoc.org/github.com/chai2010/assert.go)


Install
=======

1. `go get github.com/chai2010/assert.go`
2. `go test`

Example
=======

```Go
package assert_test

import (
	"math"
	"testing"

	. "github.com/chai2010/assert.go"
)

func TestAssert(t *testing.T) {
	Assert(t, 1 == 1)
}

func TestAssertEQ(t *testing.T) {
	AssertEQ(t, 2, 1+1)
}

func TestAssertNear(t *testing.T) {
	AssertNear(t, math.Sqrt(2), 1.414, 0.1)
}

func TestAssertMatch(t *testing.T) {
	AssertMatch(t, `^\w+@\w+\.com$`, "chaishushan@gmail.com")
}
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
