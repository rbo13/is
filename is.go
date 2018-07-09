package is

import (
	"reflect"
	"testing"
)

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

// Int64 returns an error if val
// is not of type int64
func (i *Is) Int64(val interface{}) {
	if reflect.TypeOf(val).Kind() != reflect.Int64 {
		i.testing.Errorf("Expecting as type int64, but got: %v", reflect.TypeOf(val).Kind())
	}
}
