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
		persona := person

		persona = append(persona, "Richard")
		persona = append(persona, "Burk")
		persona = append(persona, 10)

		assert.NoError(err)
		assert.NotNil(num)
		assert.ArrayEmpty(person)
		assert.NotEmpty(persona)

		assert.ArrayData(persona).SizeOf(3)
		// assert.Array(person).SizeOf(10)
	})
}
