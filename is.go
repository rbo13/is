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

// Int64 evaluates val if of type int64
func (i *Is) Int64(val interface{}) {
	i.assertVal(val, reflect.Int64)
}

// Int32 evaluates val if of type int32
func (i *Is) Int32(val interface{}) {
	i.assertVal(val, reflect.Int32)
}

// String evalueates val if of type string
func (i *Is) String(val interface{}) {
	i.assertVal(val, reflect.String)
}

func (i *Is) assertVal(val interface{}, kind reflect.Kind) {
	if reflect.TypeOf(val).Kind() != kind {
		i.testing.Errorf("AssertionError: expected as type %v, but got: %v instead", kind, reflect.TypeOf(val).Kind())
	}
}
