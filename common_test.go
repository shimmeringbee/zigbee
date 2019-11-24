package zigbee

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIEEEAddress_String(t *testing.T) {
	t.Run("converts ieee address to string", func(t *testing.T) {
		expectedValue := "00127aa9fe249cc4"
		value := IEEEAddress(0x00127aa9fe249cc4)

		assert.Equal(t, expectedValue, value.String())
	})
}