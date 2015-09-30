// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// testing helper, please fork this file to `xxx_assert_test.go`, and fix the package name.

package assert

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func tCallerFileLine(skip int) (file string, line int) {
	_, file, line, ok := runtime.Caller(skip + 1)
	if ok {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
		line = 1
	}
	return
}

func Assert(t testing.TB, condition bool, args ...interface{}) {
	if !condition {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: Assert failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: Assert failed", file, line)
		}
	}
}

func AssertNil(t testing.TB, p interface{}, args ...interface{}) {
	if p != nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("%s:%d: AssertNil failed, err = %v, %s", file, line, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertNil failed, %s", file, line, msg)
			}
		} else {
			if err, ok := p.(error); ok && err != nil {
				t.Fatalf("%s:%d: AssertNil failed, err = %v", file, line, err)
			} else {
				t.Fatalf("%s:%d: AssertNil failed", file, line)
			}
		}
	}
}

func AssertNotNil(t testing.TB, p interface{}, args ...interface{}) {
	if p == nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotNil failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotNil failed", file, line)
		}
	}
}

func AssertTrue(t testing.TB, condition bool, args ...interface{}) {
	if condition != true {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertTrue failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: AssertTrue failed", file, line)
		}
	}
}

func AssertFalse(t testing.TB, condition bool, args ...interface{}) {
	if condition != false {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertFalse failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: AssertFalse failed", file, line)
		}
	}
}

func AssertEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	// reflect.DeepEqual is failed for `int == int64?`
	if fmt.Sprintf("%q", expected) != fmt.Sprintf("%q", got) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertEqual failed, expected = %v, got = %v, %s", file, line, expected, got, msg)
		} else {
			t.Fatalf("%s:%d: AssertEqual failed, expected = %v, got = %v", file, line, expected, got)
		}
	}
}

func AssertNotEqual(t testing.TB, expected, got interface{}, args ...interface{}) {
	// reflect.DeepEqual is failed for `int == int64?`
	if fmt.Sprintf("%q", expected) == fmt.Sprintf("%q", got) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotEqual failed, expected = %v, got = %v, %s", file, line, expected, got, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotEqual failed, expected = %v, got = %v", file, line, expected, got)
		}
	}
}

func AssertNear(t testing.TB, expected, got, abs float64, args ...interface{}) {
	if math.Abs(expected-got) > abs {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNear failed, expected = %v, got = %v, abs = %v, %s", file, line, expected, got, abs, msg)
		} else {
			t.Fatalf("%s:%d: AssertNear failed, expected = %v, got = %v, abs = %v", file, line, expected, got, abs)
		}
	}
}

func AssertBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	if val < min || max < val {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertBetween failed, min = %v, max = %v, val = %v, %s", file, line, min, max, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertBetween failed, min = %v, max = %v, val = %v", file, line, min, max, val)
		}
	}
}

func AssertNotBetween(t testing.TB, min, max, val float64, args ...interface{}) {
	if min <= val && val <= max {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotBetween failed, min = %v, max = %v, val = %v, %s", file, line, min, max, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotBetween failed, min = %v, max = %v, val = %v", file, line, min, max, val)
		}
	}
}

func AssertMatch(t testing.TB, expectedPattern string, got []byte, args ...interface{}) {
	if matched, err := regexp.Match(expectedPattern, got); err != nil || !matched {
		file, line := tCallerFileLine(1)
		if err != nil {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: AssertMatch failed, expected = %q, got = %v, err = %v, %s", file, line, expectedPattern, got, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertMatch failed, expected = %q, got = %v, err = %v", file, line, expectedPattern, got, err)
			}
		} else {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: AssertMatch failed, expected = %q, got = %v, %s", file, line, expectedPattern, got, msg)
			} else {
				t.Fatalf("%s:%d: AssertMatch failed, expected = %q, got = %v", file, line, expectedPattern, got)
			}
		}
	}
}

func AssertMatchString(t testing.TB, expectedPattern, got string, args ...interface{}) {
	if matched, err := regexp.MatchString(expectedPattern, got); err != nil || !matched {
		file, line := tCallerFileLine(1)
		if err != nil {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: AssertMatchString failed, expected = %q, got = %v, err = %v, %s", file, line, expectedPattern, got, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertMatchString failed, expected = %q, got = %v, err = %v", file, line, expectedPattern, got, err)
			}
		} else {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("%s:%d: AssertMatchString failed, expected = %q, got = %v, %s", file, line, expectedPattern, got, msg)
			} else {
				t.Fatalf("%s:%d: AssertMatchString failed, expected = %q, got = %v", file, line, expectedPattern, got)
			}
		}
	}
}

func AssertSliceContain(t testing.TB, slice, val interface{}, args ...interface{}) {
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
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertSliceContain failed, slice = %v, val = %v, %s", file, line, slice, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertSliceContain failed, slice = %v, val = %v", file, line, slice, val)
		}
	}
}

func AssertSliceNotContain(t testing.TB, slice, val interface{}, args ...interface{}) {
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
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertSliceNotContain failed, slice = %v, val = %v, %s", file, line, slice, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertSliceNotContain failed, slice = %v, val = %v", file, line, slice, val)
		}
	}
}

func AssertMapContain(t testing.TB, m, key, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if !elemVal.IsValid() || !reflect.DeepEqual(elemVal.Interface(), val) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapContain failed, map = %v, key = %v, val = %v, %s", file, line, m, key, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapContain failed, map = %v, key = %v, val = %v", file, line, m, key, val)
		}
	}
}

func AssertMapContainKey(t testing.TB, m, key interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapContainKey called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if !elemVal.IsValid() {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapContainKey failed, map = %v, key = %v, %s", file, line, m, key, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapContainKey failed, map = %v, key = %v", file, line, m, key)
		}
	}
}

func AssertMapContainVal(t testing.TB, m, val interface{}, args ...interface{}) {
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
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapContainVal failed, map = %v, val = %v, %s", file, line, m, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapContainVal failed, map = %v, val = %v", file, line, m, val)
		}
	}
}

func AssertMapNotContain(t testing.TB, m, key, val interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapNotContain called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if elemVal.IsValid() && reflect.DeepEqual(elemVal.Interface(), val) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapNotContain failed, map = %v, key = %v, val = %v, %s", file, line, m, key, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapNotContain failed, map = %v, key = %v, val = %v", file, line, m, key, val)
		}
	}
}

func AssertMapNotContainKey(t testing.TB, m, key interface{}, args ...interface{}) {
	mapVal := reflect.ValueOf(m)
	if mapVal.Kind() != reflect.Map {
		panic(fmt.Sprintf("AssertMapNotContainKey called with non-map value of type %T", m))
	}
	elemVal := mapVal.MapIndex(reflect.ValueOf(key))
	if elemVal.IsValid() {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapNotContainKey failed, map = %v, key = %v, %s", file, line, m, key, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapNotContainKey failed, map = %v, key = %v", file, line, m, key)
		}
	}
}

func AssertMapNotContainVal(t testing.TB, m, val interface{}, args ...interface{}) {
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
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertMapNotContainVal failed, map = %v, val = %v, %s", file, line, m, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertMapNotContainVal failed, map = %v, val = %v", file, line, m, val)
		}
	}
}

func AssertZero(t testing.TB, val interface{}, args ...interface{}) {
	if !reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertZero failed, val = %v, %s", file, line, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertZero failed, val = %v", file, line, val)
		}
	}
}

func AssertNotZero(t testing.TB, val interface{}, args ...interface{}) {
	if reflect.DeepEqual(reflect.Zero(reflect.TypeOf(val)).Interface(), val) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotZero failed, val = %v, %s", file, line, val, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotZero failed, val = %v", file, line, val)
		}
	}
}

func AssertFileExists(t testing.TB, path string, args ...interface{}) {
	if _, err := os.Stat(path); err != nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			if err != nil {
				t.Fatalf("%s:%d: AssertFileExists failed, path = %v, err = %v, %s", file, line, path, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertFileExists failed, path = %v, %s", file, line, path, msg)
			}
		} else {
			if err != nil {
				t.Fatalf("%s:%d: AssertFileExists failed, path = %v, err = %v", file, line, path, err)
			} else {
				t.Fatalf("%s:%d: AssertFileExists failed, path = %v", file, line, path)
			}
		}
	}
}

func AssertFileNotExists(t testing.TB, path string, args ...interface{}) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			if err != nil {
				t.Fatalf("%s:%d: AssertFileNotExists failed, path = %v, err = %v, %s", file, line, path, err, msg)
			} else {
				t.Fatalf("%s:%d: AssertFileNotExists failed, path = %v, %s", file, line, path, msg)
			}
		} else {
			if err != nil {
				t.Fatalf("%s:%d: AssertFileNotExists failed, path = %v, err = %v", file, line, path, err)
			} else {
				t.Fatalf("%s:%d: AssertFileNotExists failed, path = %v", file, line, path)
			}
		}
	}
}

func AssertImplements(t testing.TB, interfaceObj, obj interface{}, args ...interface{}) {
	if !reflect.TypeOf(obj).Implements(reflect.TypeOf(interfaceObj).Elem()) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertImplements failed, interface = %T, obj = %T, %s", file, line, interfaceObj, obj, msg)
		} else {
			t.Fatalf("%s:%d: AssertImplements failed, interface = %T, obj = %T", file, line, interfaceObj, obj)
		}
	}
}

func AssertSameType(t testing.TB, expectedType interface{}, obj interface{}, args ...interface{}) {
	if !reflect.DeepEqual(reflect.TypeOf(obj), reflect.TypeOf(expectedType)) {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertSameType failed, expected = %T, obj = %T, %s", file, line, expectedType, obj, msg)
		} else {
			t.Fatalf("%s:%d: AssertSameType failed, expected = %T, obj = %T", file, line, expectedType, obj)
		}
	}
}

func AssertSameStruct(t testing.TB, expectedStruct interface{}, obj interface{}, args ...interface{}) {
	// type TypeA struct { A int, B float, C bool }
	// type TypeB struct { A int, B float, C bool }
	// AssertSameStruct(t, new(TypeA), new(TypeB))
	panic("TODO")
}

func AssertPanic(t testing.TB, f func(), args ...interface{}) {
	panicVal := func() (panicVal interface{}) {
		defer func() {
			panicVal = recover()
		}()
		f()
		return
	}()

	if panicVal == nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertPanic failed, %s", file, line, msg)
		} else {
			t.Fatalf("%s:%d: AssertPanic failed", file, line)
		}
	}
}

func AssertNotPanic(t testing.TB, f func(), args ...interface{}) {
	panicVal := func() (panicVal interface{}) {
		defer func() {
			panicVal = recover()
		}()
		f()
		return
	}()

	if panicVal != nil {
		file, line := tCallerFileLine(1)
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("%s:%d: AssertNotPanic failed, panic = %v, %s", file, line, panicVal, msg)
		} else {
			t.Fatalf("%s:%d: AssertNotPanic failed, panic = %v", file, line, panicVal)
		}
	}
}
