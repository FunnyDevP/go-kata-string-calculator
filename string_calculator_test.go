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
}
