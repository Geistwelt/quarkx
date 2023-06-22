package qlogging

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestActivateSpec(t *testing.T) {
	spec := "blockchain.network.consensus.T-PBFT=error:blockchain.ledger.block=warn:warn"

	ll := new(LoggerLevels)
	err := ll.ActivateSpec(spec)
	require.NoError(t, err)

	require.Equal(t, zapcore.WarnLevel, ll.DefaultLevel())

	require.Equal(t, zapcore.WarnLevel, ll.Level("blockchain.network.consensus"))
	require.Equal(t, zapcore.ErrorLevel, ll.Level("blockchain.network.consensus.T-PBFT.PREPARE"))
	require.True(t, ll.Enabled(zapcore.WarnLevel))
	require.False(t, ll.Enabled(zapcore.InfoLevel))
	require.Equal(t, "blockchain.ledger.block=warn:blockchain.network.consensus.T-PBFT=error:warn", ll.Spec())
}
