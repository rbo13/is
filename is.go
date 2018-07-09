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
func (i *Is) NoError(err interface{}) bool {
	if err != nil {
		i.Test.Error(err)
		return false
	}
	return true
}

// NotNil asserts if a val is not nil
func (i *Is) NotNil(val interface{}) bool {
	if val != nil {
		return true
	}

	i.Test.Error(val)
	return false
}
