- *赞助 BTC: 1Cbd6oGAUUyBi7X7MaR4np4nTmQZXVgkCW*
- *赞助 ETH: 0x623A3C3a72186A6336C79b18Ac1eD36e1c71A8a6*

----

# Assert for Go testing

[![Build Status](https://travis-ci.org/chai2010/assert.svg)](https://travis-ci.org/chai2010/assert)
[![Go Report Card](https://goreportcard.com/badge/github.com/chai2010/assert)](https://goreportcard.com/report/github.com/chai2010/assert)
[![GoDoc](https://godoc.org/github.com/chai2010/assert?status.svg)](https://godoc.org/github.com/chai2010/assert)

## Install

1. `go get -u github.com/chai2010/assert`
2. `go test`

## Example

```Go
package somepkg_test

import (
	. "github.com/chai2010/assert"
)

func TestAssert(t *testing.T) {
	Assert(t, 1 == 1)
	Assert(t, 1 == 1, "message1", "message2")
}

func TestAssertf(t *testing.T) {
	Assertf(t, 1 == 1, "%v:%v", "message1", "message2")
}
```

## BUGS

Report bugs to <chaishushan@gmail.com>.

Thanks!
