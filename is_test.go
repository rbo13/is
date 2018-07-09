package is_test

import (
	"log"
	"strconv"
	"testing"

	"github.com/whaangbuu/is"
)

func TestNoError(t *testing.T) {

	has := is.New(t)

	num, err := strconv.ParseInt("1", 10, 32)
	log.Print(num)
	has.NoError(err)
}
