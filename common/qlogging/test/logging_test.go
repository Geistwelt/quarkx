package test

import (
	"testing"
	"time"

	"github.com/geistwelt/quarkx/common/qlogging"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestLogNormal(t *testing.T) {
	qlogging.ActivateSpec("test=warn")
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

func TestREADME(t *testing.T) {
    spec := "blockchain.consensus=info:blockchain=error:warn"
    qlogging.ActivateSpec(spec)

    qxl := qlogging.MustGetLogger("blockchain.consensus.pbft")
    qxl.Debug("blockchain.consensus.pbft")  // 不会输出
    qxl.Info("blockchain.consensus.pbft")   // 正常输出

    qxl = qlogging.MustGetLogger("blockchain.p2p.switch")
    qxl.Info("blockchain.p2p.switch")       // 不会输出
    qxl.Error("blockchain.p2p.switch")      // 正常输出
}
