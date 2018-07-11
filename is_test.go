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
		var person []interface{}
		person = append(person, "Richard")
		person = append(person, "Burk")
		person = append(person, 21)

		assert.NoError(err)
		assert.NotNil(num)
		assert.Int64(int64(10))
		assert.Int32(int32(1))
		assert.String("1")
		assert.Array([2]string{"Hello", "World"})
		assert.ArrayEmpty(person)
	})
}
