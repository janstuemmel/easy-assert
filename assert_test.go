package assert

import (
	"fmt"
	"testing"
)

type obj struct {
	foo int
	bar string
}

type mockT struct {
	err func(string, ...interface{})
}

func (mt mockT) Helper() {}

func (mt mockT) Name() string {
	return "mock"
}

func (mt mockT) Errorf(format string, args ...interface{}) {
	mt.err(format, args...)
}

func Test(t *testing.T) {

	// initial impl
	mockImpl := mockT{
		err: func(format string, args ...interface{}) {},
	}

	t.Run("simple fail assertion", func(t *testing.T) {

		// given
		mock := mockImpl
		call := ""
		mock.err = func(format string, args ...interface{}) {
			call = fmt.Sprintf(format, args...)
		}
		a := New(mock)

		// when
		a.Equal(1, 2)

		// then
		if "Assertion failed for mock\n\twant:\t1\n\thave:\t2" != call {
			t.Error("assertion should fail")
		}
	})

	t.Run("object fail assertion", func(t *testing.T) {

		// given
		mock := mockImpl
		call := ""
		mock.err = func(format string, args ...interface{}) {
			call = fmt.Sprintf(format, args...)
		}
		a := New(mock)
		test := obj{1, "hello"}

		// when
		a.Equal(obj{1, "huhu"}, test)

		// then
		if "Assertion failed for mock\n\twant:\t{foo:1 bar:huhu}\n\thave:\t{foo:1 bar:hello}" != call {
			t.Error("assertion should fail")
		}
	})

	t.Run("object assertion", func(t *testing.T) {

		// given
		mock := mockImpl
		call := ""
		mock.err = func(format string, args ...interface{}) {
			call = fmt.Sprintf(format, args...)
		}
		a := New(mock)
		test := obj{1, "hello"}

		// when
		a.Equal(obj{1, "hello"}, test)

		// then
		if "" != call {
			t.Error("assertion should fail")
		}
	})

	t.Run("test non-struct Equal assertion", func(t *testing.T) {

		// given
		mock := mockImpl
		call := ""
		mock.err = func(format string, args ...interface{}) {
			call = fmt.Sprintf(format, args...)
		}
		test := obj{1, "hello"}

		// when
		Equal(mock, obj{1, "hello"}, test)

		// then
		if "" != call {
			t.Error("assertion should fail")
		}
	})

	t.Run("test non-struct Equal fail assertion", func(t *testing.T) {

		// given
		mock := mockImpl
		call := ""
		mock.err = func(format string, args ...interface{}) {
			call = fmt.Sprintf(format, args...)
		}
		test := obj{1, "hello"}

		// when
		Equal(mock, obj{1, "huhu"}, test)

		// then
		if "Assertion failed for mock\n\twant:\t{foo:1 bar:huhu}\n\thave:\t{foo:1 bar:hello}" != call {
			t.Error("assertion should fail")
		}

	})

}

func TestWithTesting(t *testing.T) {

	// uncomment this to see test failing
	t.SkipNow()

	// given
	test := obj{1, "hello"}

	// then
	Equal(t, obj{1, "hello"}, test)

}
