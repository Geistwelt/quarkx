// /**
// * Author: Xiangyu Wu
// * Date: 2023-06-22
// * From: hyperledger/fabric/common/flogging/zap.go
//  */

// package qlogging

// import (
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// 	"go.uber.org/zap/zapgrpc"
// )

// func NewZapLogger(core zapcore.Core, options ...zap.Option) *zap.Logger {
// 	return zap.New(
// 		core,
// 		append([]zap.Option{
// 			zap.AddCaller(),
// 			zap.AddStacktrace(zap.ErrorLevel),
// 		}, options...)...,
// 	)
// }

// func NewGRPCLogger(l *zap.Logger) *zapgrpc.Logger {
// 	l = l.WithOptions(
// 		zap.AddCaller(),
// 		zap.AddCallerSkip(3),
// 	)
// 	return zapgrpc.NewLogger(l, zapgrpc.WithDebug())
// }

// type QuarkXLogger struct {
// 	s *zap.SugaredLogger
// }

// func NewQuarkXLogger(l *zap.Logger, options ...zap.Option) *QuarkXLogger {
// 	return &QuarkXLogger{
// 		s: l.WithOptions(append(options, zap.AddCallerSkip(1))...).Sugar(),
// 	}
// }

// func (qxl *QuarkXLogger) Debug(args ...interface{})               { qxl.s.Debug(args...) }
// func (qxl *QuarkXLogger) Debugf(tmpl string, args ...interface{}) { qxl.s.Debugf(tmpl, args...) }
// func (qxl *QuarkXLogger) Debugw(msg string, kvs ...interface{})   { qxl.s.Debugw(msg, kvs...) }
// func (qxl *QuarkXLogger) Info(args ...interface{})                { qxl.s.Info(args...) }
// func (qxl *QuarkXLogger) Infof(tmpl string, args ...interface{})  { qxl.s.Infof(tmpl, args...) }
// func (qxl *QuarkXLogger) Infow(msg string, kvs ...interface{})    { qxl.s.Infow(msg, kvs...) }
// func (qxl *QuarkXLogger) Warn(args ...interface{})                { qxl.s.Warn(args...) }
// func (qxl *QuarkXLogger) Warnf(tmpl string, args ...interface{})  { qxl.s.Warnf(tmpl, args...) }
// func (qxl *QuarkXLogger) Warnw(msg string, kvs ...interface{})    { qxl.s.Warnw(msg, kvs...) }
// func (qxl *QuarkXLogger) Error(args ...interface{})               { qxl.s.Error(args...) }
// func (qxl *QuarkXLogger) Errorf(tmpl string, args ...interface{}) { qxl.s.Errorf(tmpl, args...) }
// func (qxl *QuarkXLogger) Errorw(msg string, kvs ...interface{})   { qxl.s.Errorw(msg, kvs...) }
// func (qxl *QuarkXLogger) Panic(args ...interface{})               { qxl.s.Panic(args...) }
// func (qxl *QuarkXLogger) Panicf(tmpl string, args ...interface{}) { qxl.s.Panicf(tmpl, args...) }
// func (qxl *QuarkXLogger) Panicw(msg string, kvs ...interface{})   { qxl.s.Panicw(msg, kvs...) }
// func (qxl *QuarkXLogger) Fatal(args ...interface{})               { qxl.s.Fatal(args...) }
// func (qxl *QuarkXLogger) Fatalf(tmpl string, args ...interface{}) { qxl.s.Fatalf(tmpl, args...) }
// func (qxl *QuarkXLogger) Fatalw(msg string, kvs ...interface{})   { qxl.s.Fatalw(msg, kvs...) }

// func (qxl *QuarkXLogger) With(args ...interface{}) *QuarkXLogger {
// 	return &QuarkXLogger{s: qxl.s.With(args...)}
// }

// func (qxl *QuarkXLogger) WithOptions(opts ...zap.Option) *QuarkXLogger {
// 	l := qxl.s.Desugar().WithOptions(opts...)
// 	return &QuarkXLogger{s: l.Sugar()}
// }

/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package qlogging

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
)

// NewZapLogger creates a zap logger around a new zap.Core. The core will use
// the provided encoder and sinks and a level enabler that is associated with
// the provided logger name. The logger that is returned will be named the same
// as the logger.
func NewZapLogger(core zapcore.Core, options ...zap.Option) *zap.Logger {
	return zap.New(
		core,
		append([]zap.Option{
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		}, options...)...,
	)
}

// NewGRPCLogger creates a grpc.Logger that delegates to a zap.Logger.
func NewGRPCLogger(l *zap.Logger) *zapgrpc.Logger {
	l = l.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)
	return zapgrpc.NewLogger(l, zapgrpc.WithDebug())
}

// NewQuarkXLogger creates a logger that delegates to the zap.SugaredLogger.
func NewQuarkXLogger(l *zap.Logger, options ...zap.Option) *QuarkXLogger {
	return &QuarkXLogger{
		s: l.WithOptions(append(options, zap.AddCallerSkip(1))...).Sugar(),
	}
}

// A QuarkXLogger is an adapter around a zap.SugaredLogger that provides
// structured logging capabilities while preserving much of the legacy logging
// behavior.
//
// The most significant difference between the QuarkXLogger and the
// zap.SugaredLogger is that methods without a formatting suffix (f or w) build
// the log entry message with fmt.Sprintln instead of fmt.Sprint. Without this
// change, arguments are not separated by spaces.
type QuarkXLogger struct{ s *zap.SugaredLogger }

func (f *QuarkXLogger) DPanic(args ...interface{})                    { f.s.DPanicf(formatArgs(args)) }
func (f *QuarkXLogger) DPanicf(template string, args ...interface{})  { f.s.DPanicf(template, args...) }
func (f *QuarkXLogger) DPanicw(msg string, kvPairs ...interface{})    { f.s.DPanicw(msg, kvPairs...) }
func (f *QuarkXLogger) Debug(args ...interface{})                     { f.s.Debugf(formatArgs(args)) }
func (f *QuarkXLogger) Debugf(template string, args ...interface{})   { f.s.Debugf(template, args...) }
func (f *QuarkXLogger) Debugw(msg string, kvPairs ...interface{})     { f.s.Debugw(msg, kvPairs...) }
func (f *QuarkXLogger) Error(args ...interface{})                     { f.s.Errorf(formatArgs(args)) }
func (f *QuarkXLogger) Errorf(template string, args ...interface{})   { f.s.Errorf(template, args...) }
func (f *QuarkXLogger) Errorw(msg string, kvPairs ...interface{})     { f.s.Errorw(msg, kvPairs...) }
func (f *QuarkXLogger) Fatal(args ...interface{})                     { f.s.Fatalf(formatArgs(args)) }
func (f *QuarkXLogger) Fatalf(template string, args ...interface{})   { f.s.Fatalf(template, args...) }
func (f *QuarkXLogger) Fatalw(msg string, kvPairs ...interface{})     { f.s.Fatalw(msg, kvPairs...) }
func (f *QuarkXLogger) Info(args ...interface{})                      { f.s.Infof(formatArgs(args)) }
func (f *QuarkXLogger) Infof(template string, args ...interface{})    { f.s.Infof(template, args...) }
func (f *QuarkXLogger) Infow(msg string, kvPairs ...interface{})      { f.s.Infow(msg, kvPairs...) }
func (f *QuarkXLogger) Panic(args ...interface{})                     { f.s.Panicf(formatArgs(args)) }
func (f *QuarkXLogger) Panicf(template string, args ...interface{})   { f.s.Panicf(template, args...) }
func (f *QuarkXLogger) Panicw(msg string, kvPairs ...interface{})     { f.s.Panicw(msg, kvPairs...) }
func (f *QuarkXLogger) Warn(args ...interface{})                      { f.s.Warnf(formatArgs(args)) }
func (f *QuarkXLogger) Warnf(template string, args ...interface{})    { f.s.Warnf(template, args...) }
func (f *QuarkXLogger) Warnw(msg string, kvPairs ...interface{})      { f.s.Warnw(msg, kvPairs...) }
func (f *QuarkXLogger) Warning(args ...interface{})                   { f.s.Warnf(formatArgs(args)) }
func (f *QuarkXLogger) Warningf(template string, args ...interface{}) { f.s.Warnf(template, args...) }

// for backwards compatibility
func (f *QuarkXLogger) Critical(args ...interface{})                   { f.s.Errorf(formatArgs(args)) }
func (f *QuarkXLogger) Criticalf(template string, args ...interface{}) { f.s.Errorf(template, args...) }
func (f *QuarkXLogger) Notice(args ...interface{})                     { f.s.Infof(formatArgs(args)) }
func (f *QuarkXLogger) Noticef(template string, args ...interface{})   { f.s.Infof(template, args...) }

func (f *QuarkXLogger) Named(name string) *QuarkXLogger { return &QuarkXLogger{s: f.s.Named(name)} }
func (f *QuarkXLogger) Sync() error                     { return f.s.Sync() }
func (f *QuarkXLogger) Zap() *zap.Logger                { return f.s.Desugar() }

func (f *QuarkXLogger) IsEnabledFor(level zapcore.Level) bool {
	return f.s.Desugar().Core().Enabled(level)
}

func (f *QuarkXLogger) With(args ...interface{}) *QuarkXLogger {
	return &QuarkXLogger{s: f.s.With(args...)}
}

func (f *QuarkXLogger) WithOptions(opts ...zap.Option) *QuarkXLogger {
	l := f.s.Desugar().WithOptions(opts...)
	return &QuarkXLogger{s: l.Sugar()}
}

func formatArgs(args []interface{}) string { return strings.TrimSuffix(fmt.Sprintln(args...), "\n") }
