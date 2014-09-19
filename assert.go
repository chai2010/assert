// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package assert provides assert helper functions for testing package.

Example:

	package assert_test

	import (
		"bytes"
		"image"
		"math"
		"strings"
		"testing"

		. "github.com/chai2010/assert.go"
	)

	func TestAssert(t *testing.T) {
		Assert(t, 1 == 1)
		Assert(t, 1 == 1, "message1", "message2")
	}

	func TestAssertTrue(t *testing.T) {
		AssertTrue(t, true)
	}

	func TestAssertFalse(t *testing.T) {
		AssertFalse(t, false)
	}

	func TestAssertEqual(t *testing.T) {
		AssertEqual(t, 2, 1+1)
		AssertEqual(t, "abc", strings.ToLower("ABC"))
		AssertEqual(t, []byte("abc"), bytes.ToLower([]byte("ABC")))
		AssertEqual(t, image.Pt(1, 2), image.Pt(1, 2))
	}

	func TestAssertNotEqual(t *testing.T) {
		AssertNotEqual(t, 2, 1)
		AssertNotEqual(t, "ABC", strings.ToLower("ABC"))
		AssertNotEqual(t, []byte("ABC"), bytes.ToLower([]byte("ABC")))
		AssertNotEqual(t, image.Pt(1, 2), image.Pt(2, 2))
		AssertNotEqual(t, image.Pt(1, 2), image.Rect(1, 2, 3, 4))
	}

	func TestAssertNear(t *testing.T) {
		AssertNear(t, 1.414, math.Sqrt(2), 0.1)
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
		AssertMatch(t, `^\w+@\w+\.com$`, []byte("chaishushan@gmail.com"))
		AssertMatch(t, `^assert`, []byte("assert.go"))
		AssertMatch(t, `\.go$`, []byte("assert.go"))
	}

	func TestAssertMatchString(t *testing.T) {
		AssertMatchString(t, `^\w+@\w+\.com$`, "chaishushan@gmail.com")
		AssertMatchString(t, `^assert`, "assert.go")
		AssertMatchString(t, `\.go$`, "assert.go")
	}

	func TestAssertSliceContain(t *testing.T) {
		AssertSliceContain(t, []int{1, 1, 2, 3, 5, 8, 13}, 8)
		AssertSliceContain(t, []interface{}{1, 1, 2, 3, 5, "8", 13}, "8")
	}

	func TestAssertSliceNotContain(t *testing.T) {
		AssertSliceNotContain(t, []int{1, 1, 2, 3, 5, 8, 13}, 12)
		AssertSliceNotContain(t, []interface{}{1, 1, 2, 3, 5, "8", 13}, 8)
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

Report bugs to <chaishushan@gmail.com>.

Thanks!
*/
package assert

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"testing"
)

func Assert(t testing.TB, condition bool, args ...interface{}) {
	if !condition {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("Assert failed, %s", msg)
		} else {
			t.Fatal("Assert failed")
		}
	}
}

func AssertTrue(t testing.TB, condition bool, args ...interface{}) {
	if condition != true {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertTrue failed, %s", msg)
		} else {
			t.Fatal("AssertTrue failed")
		}
	}
}

func AssertFalse(t testing.TB, condition bool, args ...interface{}) {
	if condition != false {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertFalse failed, %s", msg)
		} else {
			t.Fatal("AssertFalse failed")
		}
	}
}

func AssertEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	if !reflect.DeepEqual(expected, got) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertEqual failed, expected = %v, got = %v, %s", expected, got, msg)
		} else {
			t.Fatalf("AssertEqual failed, expected = %v, got = %v", expected, got)
		}
	}
}

func AssertNotEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	if reflect.DeepEqual(expected, got) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNotEqual failed, expected = %v, got = %v, %s", expected, got, msg)
		} else {
			t.Fatalf("AssertNotEqual failed, expected = %v, got = %v", expected, got)
		}
	}
}

func AssertNear(t testing.TB, expected, got, abs float64, args ...interface{}) {
	if math.Abs(expected-got) > abs {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNear failed, expected = %v, got = %v, abs = %v, %s", expected, got, abs, msg)
		} else {
			t.Fatalf("AssertNear failed, expected = %v, got = %v, abs = %v", expected, got, abs)
		}
	}
}

func AssertBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	if val < min || max < val {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertBetween failed, min = %v, max = %v, val = %v, %s", min, max, val, msg)
		} else {
			t.Fatalf("AssertBetween failed, min = %v, max = %v, val = %v", min, max, val)
		}
	}
}

func AssertNotBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	if min <= val && val <= max {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNotBetween failed, min = %v, max = %v, val = %v, %s", min, max, val, msg)
		} else {
			t.Fatalf("AssertNotBetween failed, min = %v, max = %v, val = %v", min, max, val)
		}
	}
}

func AssertMatch(t testing.TB, expectedPattern string, got []byte, args ...interface{}) {
	if matched, err := regexp.Match(expectedPattern, got); err != nil || !matched {
		if err != nil {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("AssertMatch failed, expected = %q, got = %v, err = %v, %s", expectedPattern, got, err, msg)
			} else {
				t.Fatalf("AssertMatch failed, expected = %q, got = %v, err = %v", expectedPattern, got, err)
			}
		} else {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("AssertMatch failed, expected = %q, got = %v, %s", expectedPattern, got, msg)
			} else {
				t.Fatalf("AssertMatch failed, expected = %q, got = %v", expectedPattern, got)
			}
		}
	}
}

func AssertMatchString(t testing.TB, expectedPattern, got string, args ...interface{}) {
	if matched, err := regexp.MatchString(expectedPattern, got); err != nil || !matched {
		if err != nil {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("AssertMatchString failed, expected = %q, got = %v, err = %v, %s", expectedPattern, got, err, msg)
			} else {
				t.Fatalf("AssertMatchString failed, expected = %q, got = %v, err = %v", expectedPattern, got, err)
			}
		} else {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("AssertMatchString failed, expected = %q, got = %v, %s", expectedPattern, got, msg)
			} else {
				t.Fatalf("AssertMatchString failed, expected = %q, got = %v", expectedPattern, got)
			}
		}
	}
}

func AssertSliceContain(t testing.TB, slice, elem interface{}, args ...interface{}) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic(fmt.Sprintf("AssertSliceContain called with non-slice value of type %T", slice))
	}
	var contained bool
	for i := 0; i < sliceVal.Len(); i++ {
		if reflect.DeepEqual(sliceVal.Index(i).Interface(), elem) {
			contained = true
			break
		}
	}
	if !contained {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertSliceContain failed, slice = %v, elem = %v, %s", slice, elem, msg)
		} else {
			t.Fatalf("AssertSliceContain failed, slice = %v, elem = %v", slice, elem)
		}
	}
}

func AssertSliceNotContain(t testing.TB, slice, elem interface{}, args ...interface{}) {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic(fmt.Sprintf("AssertSliceNotContain called with non-slice value of type %T", slice))
	}
	var contained bool
	for i := 0; i < sliceVal.Len(); i++ {
		if reflect.DeepEqual(sliceVal.Index(i).Interface(), elem) {
			contained = true
			break
		}
	}
	if contained {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertSliceNotContain failed, slice = %v, elem = %v, %s", slice, elem, msg)
		} else {
			t.Fatalf("AssertSliceNotContain failed, slice = %v, elem = %v", slice, elem)
		}
	}
}

func AssertMapContain(t testing.TB, m, key, elem interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if !elemVal.IsValid() && !reflect.DeepEqual(elemVal.Interface(), elem) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertMapContain failed, map = %v, key = %v, elem = %v, %s", m, key, elem, msg)
		} else {
			t.Fatalf("AssertMapContain failed, map = %v, key = %v, elem = %v", m, key, elem)
		}
	}
}

func AssertMapNotContain(t testing.TB, m, key, elem interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapNotContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if elemVal.IsValid() && reflect.DeepEqual(elemVal.Interface(), elem) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertMapNotContain failed, map = %v, key = %v, elem = %v, %s", m, key, elem, msg)
		} else {
			t.Fatalf("AssertMapNotContain failed, map = %v, key = %v, elem = %v", m, key, elem)
		}
	}
}

func AssertZero(t testing.TB, val interface{}, args ...interface{}) {
	if !reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertZero failed, val = %v, %s", val, msg)
		} else {
			t.Fatalf("AssertZero failed, val = %v", val)
		}
	}
}

func AssertNotZero(t testing.TB, val interface{}, args ...interface{}) {
	if reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNotZero failed, val = %v, %s", val, msg)
		} else {
			t.Fatalf("AssertNotZero failed, val = %v", val)
		}
	}
}

func AssertFileExists(t testing.TB, path string, args ...interface{}) {
	if _, err := os.Stat(path); err != nil {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertFileExists failed, path = %v, err = %v, %s", path, err, msg)
		} else {
			t.Fatalf("AssertFileExists failed, path = %v, err = %v", path, err)
		}
	}
}

func AssertFileNotExists(t testing.TB, path string, args ...interface{}) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertFileNotExists failed, path = %v, err = %v, %s", path, err, msg)
		} else {
			t.Fatalf("AssertFileNotExists failed, path = %v, err = %v", path, err)
		}
	}
}
