package test

import (
	"testing"
	"time"

	"github.com/geistwelt/quarkx/common/qlogging"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestLogNormal(t *testing.T) {
	// qlogging.ActivateSpec("test=debug")
	qxl := qlogging.MustGetLogger("test")
	require.NotNil(t, qxl)

	qxl.Info("hello, world")
	require.Equal(t, zapcore.InfoLevel.String(), qlogging.DefaultLevel())

	qlogging.Global.Logger("test").Error("test")
}

func TestAsyncLogging(t *testing.T) {
	qxl := qlogging.MustGetLogger("p2p")

	qxl.Info("hello")

	// var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func(i int) {
			// wg.Wait()
			qxl.Infof("message is %d", i)
		}(i)
	}

	time.Sleep(time.Second)
	// wg.Done()

}
