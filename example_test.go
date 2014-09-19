// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
