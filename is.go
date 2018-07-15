package is

import (
	"testing"
)

// Is the instance of our assertion
// package
type Is struct {
	testing *testing.T
}

// ArrayDataSet represents the array container
// and has a record of Is
type ArrayDataSet struct {
	arr    []interface{}
	assert Is
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

// ArrayEmpty evaluates if an array passed
// is empty.
func (i *Is) ArrayEmpty(arr []interface{}) {
	if len(arr) > 0 {
		i.testing.Errorf("AssertionError: expected array to be empty, but got size of: %d", len(arr))
	}
}

// NotEmpty evaluates if an array passed
// is empty.
func (i *Is) NotEmpty(arr []interface{}) {
	if arr == nil {
		i.testing.Errorf("AssertionError: expected array should not be empty, but got size of: %d", len(arr))
	}
}

// ArrayData accepts array of data
func (i Is) ArrayData(arr []interface{}) ArrayDataSet {
	return ArrayDataSet{
		arr:    arr,
		assert: i,
	}
}

// SizeOf determines the size of array
func (as ArrayDataSet) SizeOf(size int) {
	if len(as.arr) != size {
		as.assert.testing.Errorf("AssertionError: expected array to be size of: %d, but got size of %d instead", size, len(as.arr))
	}
}
