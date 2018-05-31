// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.9

package assert

import (
	"fmt"
	"image"
	"math"
	"os"
	"reflect"
	"regexp"
	"testing"
)

func Assert(t testing.TB, condition bool, args ...interface{}) {
	t.Helper()
	if !condition {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("Assert failed, %s", msg)
		} else {
			t.Fatalf("Assert failed")
		}
	}
}

func Assertf(t testing.TB, condition bool, format string, a ...interface{}) {
	t.Helper()
	if !condition {
		if msg := fmt.Sprintf(format, a...); msg != "" {
			t.Fatalf("tAssert failed, %s", msg)
		} else {
			t.Fatalf("tAssert failed")
		}
	}
}

func AssertNil(t testing.TB, p interface{}, args ...interface{}) {
	t.Helper()
	if p != nil {
		if msg := fmt.Sprint(args...); msg != "" {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("AssertNil failed, err = %v, %s", err, msg)
			} else {
				t.Fatalf("AssertNil failed, %s", msg)
			}
		} else {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("AssertNil failed, err = %v", err)
			} else {
				t.Fatalf("AssertNil failed")
			}
		}
	}
}

func AssertNotNil(t testing.TB, p interface{}, args ...interface{}) {
	t.Helper()
	if p == nil {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNotNil failed, %s", msg)
		} else {
			t.Fatalf("AssertNotNil failed")
		}
	}
}

func AssertTrue(t testing.TB, condition bool, args ...interface{}) {
	t.Helper()
	if condition != true {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertTrue failed, %s", msg)
		} else {
			t.Fatalf("AssertTrue failed")
		}
	}
}

func AssertFalse(t testing.TB, condition bool, args ...interface{}) {
	t.Helper()
	if condition != false {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertFalse failed, %s", msg)
		} else {
			t.Fatalf("AssertFalse failed")
		}
	}
}

func AssertEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	t.Helper()
	// reflect.DeepEqual is failed for `int == int64?`
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", got) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertEqual failed, expected = %v, got = %v, %s", expected, got, msg)
		} else {
			t.Fatalf("AssertEqual failed, expected = %v, got = %v", expected, got)
		}
	}
}

func AssertNotEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	// reflect.DeepEqual is failed for `int == int64?`
	if fmt.Sprintf("%v", expected) == fmt.Sprintf("%v", got) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNotEqual failed, expected = %v, got = %v, %s", expected, got, msg)
		} else {
			t.Fatalf("AssertNotEqual failed, expected = %v, got = %v", expected, got)
		}
	}
}

func AssertNear(t testing.TB, expected, got, abs float64, args ...interface{}) {
	t.Helper()
	if math.Abs(expected-got) > abs {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNear failed, expected = %v, got = %v, abs = %v, %s", expected, got, abs, msg)
		} else {
			t.Fatalf("AssertNear failed, expected = %v, got = %v, abs = %v", expected, got, abs)
		}
	}
}

func AssertBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	t.Helper()
	if val < min || max < val {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertBetween failed, min = %v, max = %v, val = %v, %s", min, max, val, msg)
		} else {
			t.Fatalf("AssertBetween failed, min = %v, max = %v, val = %v", min, max, val)
		}
	}
}

func AssertNotBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	t.Helper()
	if min <= val && val <= max {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNotBetween failed, min = %v, max = %v, val = %v, %s", min, max, val, msg)
		} else {
			t.Fatalf("AssertNotBetween failed, min = %v, max = %v, val = %v", min, max, val)
		}
	}
}

func AssertMatch(t testing.TB, expectedPattern string, got []byte, args ...interface{}) {
	t.Helper()
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
	t.Helper()
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

func AssertSliceContain(t testing.TB, slice, val interface{}, args ...interface{}) {
	t.Helper()
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic(fmt.Sprintf("AssertSliceContain called with non-slice value of type %T", slice))
	}
	var contained bool
	for i := 0; i < sliceVal.Len(); i++ {
		if reflect.DeepEqual(sliceVal.Index(i).Interface(), val) {
			contained = true
			break
		}
	}
	if !contained {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertSliceContain failed, slice = %v, val = %v, %s", slice, val, msg)
		} else {
			t.Fatalf("AssertSliceContain failed, slice = %v, val = %v", slice, val)
		}
	}
}

func AssertSliceNotContain(t testing.TB, slice, val interface{}, args ...interface{}) {
	t.Helper()
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		panic(fmt.Sprintf("AssertSliceNotContain called with non-slice value of type %T", slice))
	}
	var contained bool
	for i := 0; i < sliceVal.Len(); i++ {
		if reflect.DeepEqual(sliceVal.Index(i).Interface(), val) {
			contained = true
			break
		}
	}
	if contained {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertSliceNotContain failed, slice = %v, val = %v, %s", slice, val, msg)
		} else {
			t.Fatalf("AssertSliceNotContain failed, slice = %v, val = %v", slice, val)
		}
	}
}

func AssertMapEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	t.Helper()
	expectedMap := reflect.ValueOf(expected)
	if expectedMap.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapEqual called with non-map expected value of type %T", expected))
	}
	gotMap := reflect.ValueOf(got)
	if gotMap.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapEqual called with non-map got value of type %T", got))
	}

	if a, b := expectedMap.Len(), gotMap.Len(); a != b {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertMapEqual failed, len(expected) = %d, len(got) = %d, %s", a, b, msg)
		} else {
			t.Fatalf("AssertMapEqual failed, len(expected) = %d, len(got) = %d", a, b)
		}
		return
	}

	for _, key := range expectedMap.MapKeys() {
		expectedVal := expectedMap.MapIndex(key).Interface()
		gotVal := gotMap.MapIndex(key).Interface()

		if fmt.Sprintf("%v", expectedVal) != fmt.Sprintf("%v", gotVal) {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf(
					"AssertMapEqual failed, key = %v, expected = %v, got = %v, %s",
					key.Interface(), expectedVal, gotVal, msg,
				)
			} else {
				t.Fatalf(
					"AssertMapEqual failed, key = %v, expected = %v, got = %v",
					key.Interface(), expectedVal, gotVal,
				)
			}
			return
		}
	}
}

func AssertMapContain(t testing.TB, m, key, val interface{}, args ...interface{}) {
	t.Helper()
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if !elemVal.IsValid() || !reflect.DeepEqual(elemVal.Interface(), val) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertMapContain failed, map = %v, key = %v, val = %v, %s", m, key, val, msg)
		} else {
			t.Fatalf("AssertMapContain failed, map = %v, key = %v, val = %v", m, key, val)
		}
	}
}

func AssertMapContainKey(t testing.TB, m, key interface{}, args ...interface{}) {
	t.Helper()
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapContainKey called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if !elemVal.IsValid() {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertMapContainKey failed, map = %v, key = %v, %s", m, key, msg)
		} else {
			t.Fatalf("AssertMapContainKey failed, map = %v, key = %v", m, key)
		}
	}
}

func AssertMapContainVal(t testing.TB, m, val interface{}, args ...interface{}) {
	t.Helper()
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapContainVal called with non-map value of type %T", m))
	}
	var contained bool
	for _, key := range mapVal.MapKeys() {
		elemVal := mapVal.MapIndex(key)
		if elemVal.IsValid() && reflect.DeepEqual(elemVal.Interface(), val) {
			contained = true
			break
		}
	}
	if !contained {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertMapContainVal failed, map = %v, val = %v, %s", m, val, msg)
		} else {
			t.Fatalf("AssertMapContainVal failed, map = %v, val = %v", m, val)
		}
	}
}

func AssertMapNotContain(t testing.TB, m, key, val interface{}, args ...interface{}) {
	t.Helper()
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapNotContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if elemVal.IsValid() && reflect.DeepEqual(elemVal.Interface(), val) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertMapNotContain failed, map = %v, key = %v, val = %v, %s", m, key, val, msg)
		} else {
			t.Fatalf("AssertMapNotContain failed, map = %v, key = %v, val = %v", m, key, val)
		}
	}
}

func AssertMapNotContainKey(t testing.TB, m, key interface{}, args ...interface{}) {
	t.Helper()
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapNotContainKey called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if elemVal.IsValid() {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertMapNotContainKey failed, map = %v, key = %v, %s", m, key, msg)
		} else {
			t.Fatalf("AssertMapNotContainKey failed, map = %v, key = %v", m, key)
		}
	}
}

func AssertMapNotContainVal(t testing.TB, m, val interface{}, args ...interface{}) {
	t.Helper()
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapNotContainVal called with non-map value of type %T", m))
	}
	var contained bool
	for _, key := range mapVal.MapKeys() {
		elemVal := mapVal.MapIndex(key)
		if elemVal.IsValid() && reflect.DeepEqual(elemVal.Interface(), val) {
			contained = true
			break
		}
	}
	if contained {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertMapNotContainVal failed, map = %v, val = %v, %s", m, val, msg)
		} else {
			t.Fatalf("AssertMapNotContainVal failed, map = %v, val = %v", m, val)
		}
	}
}

func AssertZero(t testing.TB, val interface{}, args ...interface{}) {
	t.Helper()
	if !reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertZero failed, val = %v, %s", val, msg)
		} else {
			t.Fatalf("AssertZero failed, val = %v", val)
		}
	}
}

func AssertNotZero(t testing.TB, val interface{}, args ...interface{}) {
	t.Helper()
	if reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNotZero failed, val = %v, %s", val, msg)
		} else {
			t.Fatalf("AssertNotZero failed, val = %v", val)
		}
	}
}

func AssertFileExists(t testing.TB, path string, args ...interface{}) {
	t.Helper()
	if _, err := os.Stat(path); err != nil {
		if msg := fmt.Sprint(args...); msg != "" {
			if err != nil {
				t.Fatalf("AssertFileExists failed, path = %v, err = %v, %s", path, err, msg)
			} else {
				t.Fatalf("AssertFileExists failed, path = %v, %s", path, msg)
			}
		} else {
			if err != nil {
				t.Fatalf("AssertFileExists failed, path = %v, err = %v", path, err)
			} else {
				t.Fatalf("AssertFileExists failed, path = %v", path)
			}
		}
	}
}

func AssertFileNotExists(t testing.TB, path string, args ...interface{}) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if msg := fmt.Sprint(args...); msg != "" {
			if err != nil {
				t.Fatalf("AssertFileNotExists failed, path = %v, err = %v, %s", path, err, msg)
			} else {
				t.Fatalf("AssertFileNotExists failed, path = %v, %s", path, msg)
			}
		} else {
			if err != nil {
				t.Fatalf("AssertFileNotExists failed, path = %v, err = %v", path, err)
			} else {
				t.Fatalf("AssertFileNotExists failed, path = %v", path)
			}
		}
	}
}

func AssertImplements(t testing.TB, interfaceObj, obj interface{}, args ...interface{}) {
	t.Helper()
	if !reflect.TypeOf(obj).Implements(reflect.TypeOf(interfaceObj).Elem()) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertImplements failed, interface = %T, obj = %T, %s", interfaceObj, obj, msg)
		} else {
			t.Fatalf("AssertImplements failed, interface = %T, obj = %T", interfaceObj, obj)
		}
	}
}

func AssertSameType(t testing.TB, expectedType interface{}, obj interface{}, args ...interface{}) {
	t.Helper()
	if !reflect.DeepEqual(reflect.TypeOf(obj), reflect.TypeOf(expectedType)) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertSameType failed, expected = %T, obj = %T, %s", expectedType, obj, msg)
		} else {
			t.Fatalf("AssertSameType failed, expected = %T, obj = %T", expectedType, obj)
		}
	}
}

func AssertPanic(t testing.TB, f func(), args ...interface{}) {
	t.Helper()

	panicVal := func() (panicVal interface{}) {
		defer func() {
			panicVal = recover()
		}()
		f()
		return
	}()

	if panicVal == nil {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertPanic failed, %s", msg)
		} else {
			t.Fatalf("AssertPanic failed")
		}
	}
}

func AssertNotPanic(t testing.TB, f func(), args ...interface{}) {
	t.Helper()

	panicVal := func() (panicVal interface{}) {
		defer func() {
			panicVal = recover()
		}()
		f()
		return
	}()

	if panicVal != nil {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNotPanic failed, panic = %v, %s", panicVal, msg)
		} else {
			t.Fatalf("AssertNotPanic failed, panic = %v", panicVal)
		}
	}
}

func AssertImageEqual(t testing.TB, expected, got image.Image, maxDelta int, args ...interface{}) {
	t.Helper()

	if equal, pos := tImageEqual(expected, got, maxDelta); !equal {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertImageEqual failed, pos = %v, expected = %v, got = %v, %s", pos, expected, got, msg)
		} else {
			t.Fatalf("AssertImageEqual failed, pos = %v, expected = %v, got = %v", pos, expected, got)
		}
	}
}

func tImageEqual(m0, m1 image.Image, maxDelta int) (ok bool, failedPixelPos image.Point) {
	b := m0.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c0 := m0.At(x, y)
			c1 := m1.At(x, y)
			r0, g0, b0, a0 := c0.RGBA()
			r1, g1, b1, a1 := c1.RGBA()
			if tDeltaInt(int(r0), int(r1)) > maxDelta {
				return false, image.Pt(x, y)
			}
			if tDeltaInt(int(g0), int(g1)) > maxDelta {
				return false, image.Pt(x, y)
			}
			if tDeltaInt(int(b0), int(b1)) > maxDelta {
				return false, image.Pt(x, y)
			}
			if tDeltaInt(int(a0), int(a1)) > maxDelta {
				return false, image.Pt(x, y)
			}
		}
	}
	return true, image.Pt(0, 0)
}

func tDeltaInt(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
