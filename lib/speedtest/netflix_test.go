package speedtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNetflix(t *testing.T) {
	result, err := GetResult(NetflixProvider)
	require.NoError(t, err)
	require.NotNil(t, result)
}