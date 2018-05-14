// +build !windows

package globalenv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvironUnixInternals(t *testing.T) {
	require := require.New(t)
	_, err := DefaultHandler.execEnv()
	require.NoError(err)

}
