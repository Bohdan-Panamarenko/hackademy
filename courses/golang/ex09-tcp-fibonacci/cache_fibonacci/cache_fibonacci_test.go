package cache_fibonacci

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func bigIntNil() *big.Int {
	return nil
}

func TestFibonacciCache(t *testing.T) {
	t.Run("create new cache", func(t *testing.T) {
		assert.EqualValues(
			t,
			&Cache{
				cleanupInterval: 0,
				items:           make(map[int64]*big.Int),
			},
			New(0),
		)
	})

	t.Run("set new field to cache", func(t *testing.T) {
		cache := New(0)

		expected := big.NewInt(5)

		cache.Set(5, expected)

		assert.EqualValues(t, expected, cache.items[5])
	})

	t.Run("get field from cache", func(t *testing.T) {
		cache := New(0)

		expected := big.NewInt(5)

		cache.Set(5, expected)

		assert.EqualValues(t, expected, cache.Get(5))
	})

	t.Run("delete field from cache", func(t *testing.T) {
		cache := New(0)

		expected := big.NewInt(5)

		cache.Set(5, expected)

		cache.Delete(5)

		assert.EqualValues(t, bigIntNil(), cache.Get(5))
	})

	t.Run("cleanup", func(t *testing.T) {
		cache := New(2 * time.Second)

		expected := big.NewInt(5)

		cache.Set(5, expected)

		time.Sleep(3 * time.Second)

		assert.EqualValues(t, bigIntNil(), cache.Get(5))
	})
}
