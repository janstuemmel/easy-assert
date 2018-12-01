# easy-assert [![Build Status](https://travis-ci.org/janstuemmel/easy-assert.svg?branch=master)](https://travis-ci.org/janstuemmel/easy-assert) [![Godoc](https://godoc.org/github.com/janstuemmel/easy-assert?status.svg)](http://godoc.org/github.com/janstuemmel/easy-assert)

A simple go assertion library. 

You don't need this library, instead just copy following lines into your test file:

```go
func assert(t *testing.T, want interface{}, have interface{}) {

  // mark as test helper
  t.Helper()

  // throw error
  if want != have {
    t.Errorf("Assertion failed for %s\n\twant:\t%+v\n\thave:\t%+v", t.Name(), want, have)
  }
}
```

## Usage

```go
package foo_test

import (
  "testing"
  "errors"

  "github.com/janstuemmel/easy-assert"
)

func TestFoo(t *testing.T) {

  assert.Equal(t, 1, 1)
  assert.Equal(t, 1, 2)

  // test thrown error
  err := errors.New("foo")
  assert.Equal(t, false, nil == err)

  // alternative
  a := assert.New(t)
  a.Equal(1, 1)
  a.Equal(1, 2)
}
```