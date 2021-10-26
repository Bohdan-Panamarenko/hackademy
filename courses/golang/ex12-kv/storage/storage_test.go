package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	test := assert.New(t)

	t.Run("test add", func(t *testing.T) {
		s := NewStorage()

		expected := "juice"

		err := s.Add("apple", expected)
		test.NoError(err)

		test.EqualValues(expected, s.kv["apple"])

		err = s.Add("", expected)
		test.EqualError(err, "key can not be empty")

		err = s.Add("orange", "")
		test.EqualError(err, "value can not be empty")

		err = s.Add("apple", expected)
		test.EqualError(err, "key 'apple' already exists")
	})

	t.Run("test get", func(t *testing.T) {
		s := NewStorage()

		expected := "juice"

		s.Add("apple", expected)

		value, err := s.Get("apple")
		test.NoError(err)

		test.EqualValues(expected, value)

		_, err = s.Get("orange")
		test.EqualError(err, "key 'orange' does not exist")
	})

	t.Run("test update", func(t *testing.T) {
		s := NewStorage()

		expected := "juice"

		s.Add("apple", expected)

		err := s.Update("", expected)
		test.EqualError(err, "key can not be empty")

		err = s.Update("apple", "")
		test.EqualError(err, "value can not be empty")

		err = s.Update("orange", expected)
		test.EqualError(err, "key 'orange' does not exist")

		newExpected := "cake"

		err = s.Update("apple", newExpected)
		test.NoError(err)
		test.Equal(newExpected, s.kv["apple"])
	})

	t.Run("test delete", func(t *testing.T) {
		s := NewStorage()

		expected := "juice"

		s.Add("apple", expected)

		value, err := s.Delete("")
		test.EqualError(err, "key can not be empty")
		test.Empty(value)

		// err = s.Update("apple", "")
		// test.EqualError(err, "value can not be empty")

		// err = s.Update("orange", expected)
		// test.EqualError(err, "key 'orange' does not exist")

		// newExpected := "cake"

		value, err = s.Delete("apple")
		test.NoError(err)
		test.Empty(s.kv["apple"])
		test.Equal(expected, value)
	})
}
