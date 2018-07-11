package is_test

import (
	"strconv"
	"testing"

	"github.com/whaangbuu/is"
)

func TestIsPackage(t *testing.T) {
	t.Run("IsPackage", func(mt *testing.T) {
		assert := is.New(mt)

		num, err := strconv.ParseInt("10", 10, 32)

		assert.NoError(err)
		assert.NotNil(num)
		assert.Int64(int32(10))
		assert.Int32(int64(1))
		assert.String(1)
	})
}
