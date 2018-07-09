package is

import "testing"

// Is the instance of our assertion
// package
type Is struct {
	testing *testing.T
}

// New returns an instance of Is
func New(t *testing.T) Is {
	return Is{
		testing: t,
	}
}

// NoError asserts if no error from
// a certain value
func (i *Is) NoError(err interface{}) bool {
	if err != nil {
		i.testing.Errorf("Error due to: %v", err)
		return i.testing.Failed()
	}
	return true
}

// NotNil asserts if a val is not nil
func (i *Is) NotNil(val interface{}) bool {
	if val != nil {
		return true
	}

	i.testing.Errorf("Value should not be empty, instead got: %v", val)
	return i.testing.Failed()
}
