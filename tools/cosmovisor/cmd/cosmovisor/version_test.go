package main

import (
	"context"
	"os"
	"testing"

	"cosmossdk.io/tools/cosmovisor"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/stretchr/testify/require"
)

func TestVersionCommand_Error(t *testing.T) {
	// Unset the environment variable
	err := os.Unsetenv("DAEMON_NAME")
	require.NoError(t, err)

	logger := cosmovisor.NewLogger()

	rootCmd.SetArgs([]string{"version"})
	_, out := testutil.ApplyMockIO(rootCmd)
	ctx := context.WithValue(context.Background(), cosmovisor.LoggerKey, logger)

	require.Error(t, rootCmd.ExecuteContext(ctx))
	require.Contains(t, out.String(), "DAEMON_NAME is not set")
}
