/**
* Author: Xiangyu Wu
* Date: 2023-06-22
* From: hyperledger/fabric/common/qlogging/core.go
 */

package qlogging

import (
	"go.uber.org/zap/zapcore"
)

type Encoding uint8

const (
	CONSOLE = 0
	JSON    = 1
	LOGFMT  = 2
)

type EncodingSelector interface {
	Encoding() Encoding
}

// Core 是一个自定义的 zapcore.Core 的实现。
type Core struct {
	zapcore.LevelEnabler
	Levels   *LoggerLevels
	Encoders map[Encoding]zapcore.Encoder
	Selector EncodingSelector
	Output   zapcore.WriteSyncer
	Observer Observer
}

type Observer interface {
	Check(e zapcore.Entry, ce *zapcore.CheckedEntry)
	WriteEntry(e zapcore.Entry, fields []zapcore.Field)
}

func (c *Core) With(fields []zapcore.Field) zapcore.Core {
	clones := map[Encoding]zapcore.Encoder{}
	for name, enc := range c.Encoders {
		clone := enc.Clone()
		addFields(clone, fields)
		clones[name] = clone
	}

	return &Core{
		LevelEnabler: c.LevelEnabler,
		Levels:       c.Levels,
		Encoders:     clones,
		Selector:     c.Selector,
		Output:       c.Output,
		Observer:     c.Observer,
	}
}

func (c *Core) Check(e zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Observer != nil {
		c.Observer.Check(e, ce)
	}

	if c.Enabled(e.Level) && c.Levels.Level(e.LoggerName).Enabled(e.Level) {
		return ce.AddCore(e, c)
	}
	return ce
}

func (c *Core) Write(e zapcore.Entry, fields []zapcore.Field) error {
	encoding := c.Selector.Encoding()
	enc := c.Encoders[encoding]

	buf, err := enc.EncodeEntry(e, fields)
	if err != nil {
		return err
	}

	_, err = c.Output.Write(buf.Bytes())
	buf.Free()
	if err != nil {
		return err
	}

	if e.Level >= zapcore.PanicLevel {
		c.Sync()
	}

	if c.Observer != nil {
		c.Observer.WriteEntry(e, fields)
	}

	return nil
}

func (c *Core) Sync() error {
	return c.Output.Sync()
}

func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}
