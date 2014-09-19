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

func TestAssertTrue(t *testing.T) {
	AssertTrue(t, true)
}

func TestAssertFalse(t *testing.T) {
	AssertFalse(t, false)
}

func TestAssertEqual(t *testing.T) {
	AssertEqual(t, 2, 1+1)
}

func TestAssertNotEqual(t *testing.T) {
	AssertNotEqual(t, 2, 1)
}

func TestAssertNear(t *testing.T) {
	AssertNear(t, math.Sqrt(2), 1.414, 0.1)
}

func TestAssertBetween(t *testing.T) {
	AssertBetween(t, 0, 255, 0)
	AssertBetween(t, 0, 255, 128)
	AssertBetween(t, 0, 255, 255)
}

func TestAssertNotBetween(t *testing.T) {
	AssertNotBetween(t, 0, 255, -1)
	AssertNotBetween(t, 0, 255, 256)
}

func TestAssertMatch(t *testing.T) {
	AssertMatch(t, `^\w+@\w+\.com$`, "chaishushan@gmail.com")
}

func TestAssertSliceContain(t *testing.T) {
	AssertSliceContain(t, []int{1, 1, 2, 3, 5, 8, 13}, 8)
}

func TestAssertSliceNotContain(t *testing.T) {
	AssertSliceNotContain(t, []int{1, 1, 2, 3, 5, 8, 13}, 12)
}

func TestAssertMapContain(t *testing.T) {
	AssertMapContain(t,
		map[string]int{
			"UTC": 0 * 60 * 60,
			"EST": -5 * 60 * 60,
			"CST": -6 * 60 * 60,
			"MST": -7 * 60 * 60,
			"PST": -8 * 60 * 60,
		},
		"MST", -7*60*60,
	)

}

func TestAssertMapNotContain(t *testing.T) {
	AssertMapNotContain(t,
		map[string]int{
			"UTC": 0 * 60 * 60,
			"EST": -5 * 60 * 60,
			"CST": -6 * 60 * 60,
			"MST": -7 * 60 * 60,
			"PST": -8 * 60 * 60,
		},
		"ABC", -7*60*60,
	)
}

func TestAssertZero(t *testing.T) {
	AssertZero(t, struct {
		A bool
		B string
		C int
		d map[string]interface{}
	}{})
}

func TestAssertNotZero(t *testing.T) {
	AssertNotZero(t, struct {
		A bool
		B string
		C int
		d map[string]interface{}
	}{A: true})
}

func TestAssertFileExists(t *testing.T) {
	AssertFileExists(t, "assert.go")
}

func TestAssertFileNotExists(t *testing.T) {
	AssertFileNotExists(t, "assert.cc")
}
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
