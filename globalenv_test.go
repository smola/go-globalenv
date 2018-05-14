package globalenv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSupported(t *testing.T) {
	require := require.New(t)
	require.True(Supported())
}

func TestEnviron(t *testing.T) {
	require := require.New(t)
	env := Environ()
	for _, e := range env {
		t.Log(e)
	}

	require.NotEmpty(env)
	require.True(Supported())
}

func TestGetenv(t *testing.T) {
	require := require.New(t)
	require.NotEmpty(Getenv("PATH"))
	require.Empty(Getenv("NON_EXISTING_ENVIRONMENT_VARIABLE"))
}
