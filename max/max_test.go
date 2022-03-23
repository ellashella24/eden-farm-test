package max

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumber(t *testing.T) {
	t.Run("Test Max Number Function", func(t *testing.T) {
		res1 := MaxNumber([]int{1, 2, 3, 8, 9, 3, 2, 1})
		res2 := MaxNumber([]int{1, 2, 1, 2, 2, 1})
		res3 := MaxNumber([]int{7, 1, 2, 9, 7, 2, 1})

		assert.Equal(t, 3, res1)
		assert.Equal(t, 2, res2)
		assert.Equal(t, 2, res3)
	})
}
