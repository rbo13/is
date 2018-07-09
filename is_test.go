package is_test

import (
	"strconv"
	"testing"

	"github.com/whaangbuu/is"
)

func TestNoError(t *testing.T) {

	has := is.New(t)

	num, err := strconv.ParseInt("1", 10, 32)

	has.NoError(err) // false
	has.NotNil(num)  // 1
}
