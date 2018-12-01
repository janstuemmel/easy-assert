package assert

type (
	t interface {
		Helper()
		Name() string
		Errorf(string, ...interface{})
	}

	// Assert struct
	Assert struct {
		t t
	}
)

var format = "Assertion failed for %s\n\twant:\t%+v\n\thave:\t%+v"

// New initializes new assert object
func New(t t) Assert {
	return Assert{t}
}

// Equal asserts that two given values are equal
func (a Assert) Equal(want interface{}, have interface{}) {
	equal(a.t, want, have)
}

// Equal asserts that two given values are equal
func Equal(t t, want interface{}, have interface{}) {
	equal(t, want, have)
}

func equal(t t, want interface{}, have interface{}) {
	t.Helper()

	if want != have {
		t.Errorf(format, t.Name(), want, have)
	}
}
