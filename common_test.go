package zigbee

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClusterID(t *testing.T) {
	t.Run("that cluster IDs over 0xfc00 are manufacturer specific", func(t *testing.T) {
		assert.False(t, ClusterID(0xfbff).IsManufacturerSpecific())
		assert.True(t, ClusterID(0xfc00).IsManufacturerSpecific())
	})
}

func TestIEEEAddress_String(t *testing.T) {
	t.Run("converts ieee address to string", func(t *testing.T) {
		expectedValue := "00127aa9fe249cc4"
		value := IEEEAddress(0x00127aa9fe249cc4)

		assert.Equal(t, expectedValue, value.String())
	})
}
