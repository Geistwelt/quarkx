/**
* Author: Xiangyu Wu
* Date: 2023-06-22
* From: hyperledger/fabric/common/flogging/metrics/observer.go
 */

package metrics

import (
	"github.com/geistwelt/quarkx/common/metrics"
	"go.uber.org/zap/zapcore"
)

var CheckedCountOpt = metrics.CounterOpts{
	Namespace:    "logging",
	Name:         "entries_checked",
	Help:         "Number of log entries checked against the active logging level",
	LabelNames:   []string{"level"},
	StatsdFormat: "%{#fqname}.%{level}",
}

var WrittenCountOpt = metrics.CounterOpts{
	Namespace:    "logging",
	Name:         "entries_written",
	Help:         "Number of log entries that are written",
	LabelNames:   []string{"level"},
	StatsdFormat: "%{#fqname}.%{level}",
}

type Observer struct {
	CheckedCounter metrics.Counter
	WrittenCounter metrics.Counter
}

func NewObserver(provider metrics.Provider) *Observer {
	return &Observer{
		CheckedCounter: provider.NewCounter(CheckedCountOpt),
		WrittenCounter: provider.NewCounter(WrittenCountOpt),
	}
}

func (o *Observer) Check(e zapcore.Entry, ce *zapcore.CheckedEntry) {
	o.CheckedCounter.With("level", e.Level.String()).Add(1)
}

func (o *Observer) WriteEntry(e zapcore.Entry, fields []zapcore.Field) {
	o.WrittenCounter.With("level", e.Level.String()).Add(1)
}
