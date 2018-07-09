package is_test

import (
	"strconv"
	"testing"

	"github.com/whaangbuu/is"
)

func TestIsPackage(t *testing.T) {
	t.Run("IsPackage", func(mt *testing.T) {
		has := is.New(mt)

		num, err := strconv.ParseInt("10", 10, 64)

		has.NoError(err)
		has.NotNil(num)
	})
}
