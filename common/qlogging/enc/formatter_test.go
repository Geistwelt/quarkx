package enc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestEncodeEntry(t *testing.T) {
	fields := []zapcore.Field{
		{
			Key:       "Name",
			Type:      0,
			Integer:   1,
			String:    "Koi",
			Interface: nil,
		},
		{
			Key:       "Age",
			Type:      0,
			Integer:   2,
			String:    "18",
			Interface: nil,
		},
	}

	encoder := NewFormatEncoder()
	entry := zapcore.Entry{
		Level:      0,
		Time:       time.Now(),
		LoggerName: "QuarkX",
		Message:    "hello, world",
		Caller:     zapcore.EntryCaller{},
		Stack:      "2",
	}

	line, err := encoder.EncodeEntry(entry, fields)
	require.NoError(t, nil, err)

	t.Log(line.String())
}
