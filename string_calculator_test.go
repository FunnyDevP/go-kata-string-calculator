package string_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringCalculator(t *testing.T) {
	t.Run("should return 0 when take empty string", func(t *testing.T) {
		result, err := Add("")

		assert.Nil(t, err)
		assert.Equal(t, 0,result)
	})
	t.Run("should return 1 when take `1` string", func(t *testing.T) {
		result, err := Add("1")

		assert.Nil(t, err)
		assert.Equal(t, 1,result)
	})
	t.Run("should return 2 when take `2` string", func(t *testing.T) {
		result,err := Add("2")

		assert.Nil(t, err)
		assert.Equal(t, 2,result)
	})
	t.Run("should return 3 when take `1,2` string", func(t *testing.T) {
		result, err := Add("1,2")

		assert.Nil(t, err)
		assert.Equal(t, 3,result)
	})
	t.Run("should return 55 when take `1,2,3,4,5,6,7,8,9,10`", func(t *testing.T) {
		result, err := Add("1,2,3,4,5,6,7,8,9,10")

		assert.Nil(t, err)
		assert.Equal(t, 55,result)
	})
	t.Run("should return 6 when take `1\n2,3`", func(t *testing.T) {
		result, err := Add("1\n2,3")

		assert.Nil(t, err)
		assert.Equal(t, 6,result)
	})
	t.Run("should return 3 when take `//;\n1;2`", func(t *testing.T) {
		result, err := Add("//;\n1;2")

		assert.Nil(t, err)
		assert.Equal(t, 3,result)
	})
	t.Run("should return error negative now allowed when take `1,-2,-3`", func(t *testing.T) {
		result, err := Add("1,-2,-3")

		assert.Zero(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "error: negatives not allowed: -2 -3",err.Error())
	})
	t.Run("should return 2 when take `1001,2`", func(t *testing.T) {
		result, err := Add("1001,2")

		assert.Nil(t, err)
		assert.Equal(t, 2,result)
	})
}
