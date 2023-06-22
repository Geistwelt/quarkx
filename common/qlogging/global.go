/**
* Author: Xiangyu Wu
* Date: 2023-06-22
* From: hyperledger/fabric/common/flogging/global.go
 */

package qlogging

import "go.uber.org/zap/zapcore"

const (
	defaultLevel = zapcore.InfoLevel
	defaultFormat = "%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{sequence:03x}%{color:reset} %{message}"
)