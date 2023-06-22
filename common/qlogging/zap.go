/**
* Author: Xiangyu Wu
* Date: 2023-06-22
* From: hyperledger/fabric/common/flogging/zap.go
 */

package qlogging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
)

func NewZapLogger(core zapcore.Core, options ...zap.Option) *zap.Logger {
	return zap.New(
		core,
		append([]zap.Option{
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
		}, options...)...,
	)
}

func NewGRPCLogger(l *zap.Logger) *zapgrpc.Logger {
	l = l.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)
	return zapgrpc.NewLogger(l, zapgrpc.WithDebug())
}

type QuarkXLogger struct {
	s *zap.SugaredLogger
}

func NewQuarkXLogger(l *zap.Logger, options ...zap.Option) *QuarkXLogger {
	return &QuarkXLogger{
		s: l.WithOptions(append(options, zap.AddCallerSkip(1))...).Sugar(),
	}
}

func (qxl *QuarkXLogger) Debug(args ...interface{})               { qxl.s.Debug(args...) }
func (qxl *QuarkXLogger) Debugf(tmpl string, args ...interface{}) { qxl.s.Debugf(tmpl, args...) }
func (qxl *QuarkXLogger) Debugw(msg string, kvs ...interface{})   { qxl.s.Debugw(msg, kvs...) }
func (qxl *QuarkXLogger) Info(args ...interface{})                { qxl.s.Info(args...) }
func (qxl *QuarkXLogger) Infof(tmpl string, args ...interface{})  { qxl.s.Infof(tmpl, args...) }
func (qxl *QuarkXLogger) Infow(msg string, kvs ...interface{})    { qxl.s.Infow(msg, kvs...) }
func (qxl *QuarkXLogger) Warn(args ...interface{})                { qxl.s.Warn(args...) }
func (qxl *QuarkXLogger) Warnf(tmpl string, args ...interface{})  { qxl.s.Warnf(tmpl, args...) }
func (qxl *QuarkXLogger) Warnw(msg string, kvs ...interface{})    { qxl.s.Warnw(msg, kvs...) }
func (qxl *QuarkXLogger) Error(args ...interface{})               { qxl.s.Error(args...) }
func (qxl *QuarkXLogger) Errorf(tmpl string, args ...interface{}) { qxl.s.Errorf(tmpl, args...) }
func (qxl *QuarkXLogger) Errorw(msg string, kvs ...interface{})   { qxl.s.Errorw(msg, kvs...) }
func (qxl *QuarkXLogger) Panic(args ...interface{})               { qxl.s.Panic(args...) }
func (qxl *QuarkXLogger) Panicf(tmpl string, args ...interface{}) { qxl.s.Panicf(tmpl, args...) }
func (qxl *QuarkXLogger) Panicw(msg string, kvs ...interface{})   { qxl.s.Panicw(msg, kvs...) }
func (qxl *QuarkXLogger) Fatal(args ...interface{})               { qxl.s.Fatal(args...) }
func (qxl *QuarkXLogger) Fatalf(tmpl string, args ...interface{}) { qxl.s.Fatalf(tmpl, args...) }
func (qxl *QuarkXLogger) Fatalw(msg string, kvs ...interface{})   { qxl.s.Fatalw(msg, kvs...) }

func (qxl *QuarkXLogger) With(args ...interface{}) *QuarkXLogger {
	return &QuarkXLogger{s: qxl.s.With(args...)}
}

func (qxl *QuarkXLogger) WithOptions(opts ...zap.Option) *QuarkXLogger {
	l := qxl.s.Desugar().WithOptions(opts...)
	return &QuarkXLogger{s: l.Sugar()}
}