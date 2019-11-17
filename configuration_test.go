package zigbee

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfiguration(t *testing.T) {
	t.Run("verify random creation", func(t *testing.T) {
		null := NetworkConfiguration{}

		actual, err := GenerateNetworkConfiguration()

		assert.NoError(t, err)
		assert.NotEqual(t, null.PANID, actual.PANID)
		assert.NotEqual(t, null.ExtendedPANID, actual.ExtendedPANID)
		assert.NotEqual(t, null.NetworkKey, actual.NetworkKey)
		assert.Equal(t, DefaultChannel, actual.Channel)
	})
}
