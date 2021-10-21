package calc_fibonacci

import (
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcFibonacci(t *testing.T) {
	t.Run("wrong length of row", func(t *testing.T) {
		resp, err := Calc(0)

		expected := func() *big.Int {
			return nil
		}()

		assert.Equal(t, expected, resp)
		assert.Equal(t, errors.New("argument can not be less than 1"), err)
	})

	t.Run("right length of row", func(t *testing.T) {
		resp, err := Calc(5)

		expected := big.NewInt(5)

		assert.EqualValues(t, *expected, *resp)
		assert.Equal(t, nil, err)
	})
}
