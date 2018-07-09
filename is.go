package is

import "testing"

// Is the instance of our assertion
// package
type Is struct {
	Test *testing.T
}

// New returns an instance of Is
func New(t *testing.T) Is {
	return Is{
		Test: t,
	}
}

// NoError asserts if no error from
// a certain value
func (is *Is) NoError(err interface{}) bool {
	if err != nil {
		return false
	}
	return true
}
