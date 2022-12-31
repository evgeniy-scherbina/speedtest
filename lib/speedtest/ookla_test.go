package speedtest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOokla(t *testing.T) {
	result, err := GetResult(OoklaProvider)
	require.NoError(t, err)
	require.NotNil(t, result)

	fmt.Printf("Result: %v\n", result)
}
