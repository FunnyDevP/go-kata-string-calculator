package string_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringCalculator(t *testing.T) {
	t.Run("should return 0 when take empty string", func(t *testing.T) {
		assert.Equal(t, 0,Add(""))
	})
	t.Run("should return 1 when take `1` string", func(t *testing.T) {
		assert.Equal(t, 1,Add("1"))
	})
	t.Run("should return 2 when take `2` string", func(t *testing.T) {
		assert.Equal(t, 2,Add("2"))
	})
	t.Run("should return 3 when take `1,2` string", func(t *testing.T) {
		assert.Equal(t, 3,Add("1,2"))
	})
	t.Run("should return 55 when take `1,2,3,4,5,6,7,8,9,10`", func(t *testing.T) {
		assert.Equal(t, 55,Add("1,2,3,4,5,6,7,8,9,10"))
	})
	t.Run("should return 6 when take `1\n2,3`", func(t *testing.T) {
		assert.Equal(t, 6,Add("1\n2,3"))
	})
	t.Run("should return 3 when take `//;\n1;2`", func(t *testing.T) {

		assert.Equal(t, 3,Add("//;\n1;2"))
	})
}
