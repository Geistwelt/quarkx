/**
* Author: Xiangyu Wu
* Date: 2023-06-22
* From: hyperledger/fabric/common/flogging/loggerlevels.go
 */

package qlogging

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"

	"go.uber.org/zap/zapcore"
)

var loggerNameRegexp = regexp.MustCompile(`^[[:alnum:]_-]+(\.[[:alnum:]_-]+)*$`)

type LoggerLevels struct {
	mutex        sync.RWMutex
	levelCache   map[string]zapcore.Level
	specs        map[string]zapcore.Level
	defaultLevel zapcore.Level
	minLevel     zapcore.Level
}

func (ll *LoggerLevels) DefaultLevel() zapcore.Level {
	ll.mutex.RLock()
	lvl := ll.defaultLevel
	ll.mutex.RUnlock()
	return lvl
}

func (ll *LoggerLevels) ActivateSpec(spec string) error {
	ll.mutex.Lock()
	defer ll.mutex.Unlock()

	defaultLevel := zapcore.InfoLevel

	specs := make(map[string]zapcore.Level)

	for _, field := range strings.Split(spec, ":") {
		split := strings.Split(field, "=")

		switch len(split) {
		case 1:
			// 不含有 =
			if field != "" && !IsValidLevel(field) {
				return fmt.Errorf("invalid logging specification '%s': bad segment '%s'", spec, field)
			}
			defaultLevel = NameToLevel(field)
		case 2:
			// 含有 =
			if split[0] == "" {
				return fmt.Errorf("invalid logging specification '%s': no logger specified in segment '%s'", spec, field)
			}
			if split[1] != "" && !IsValidLevel(split[1]) {
				return fmt.Errorf("invalid logging specification '%s': bad segment '%s'", spec, field)
			}

			level := NameToLevel(split[1])
			for _, logger := range strings.Split(split[0], ",") {
				if !isValidLoggerName(strings.TrimSuffix(logger, ".")) {
					return fmt.Errorf("invalid logging specification '%s': bad logger name '%s'", spec, logger)
				} 
				specs[logger] = level
			}
		default:
			return fmt.Errorf("invalid logging specification '%s': bad segment '%s'", spec, field)
		}
	}

	minLevel := defaultLevel
	for _, lvl := range specs {
		if minLevel > lvl {
			minLevel = lvl
		}
	}

	ll.minLevel = minLevel
	ll.defaultLevel = defaultLevel
	ll.specs = specs
	ll.levelCache = make(map[string]zapcore.Level)

	return nil
}

func (ll *LoggerLevels) Level(loggerName string) zapcore.Level {
	if lvl, ok := ll.cachedLevel(loggerName); ok {
		return lvl
	}

	ll.mutex.Lock()
	lvl := ll.calculateLevel(loggerName)
	ll.levelCache[loggerName] = lvl
	ll.mutex.Unlock()
	return lvl
}

func (ll *LoggerLevels) Spec() string {
	ll.mutex.RLock()
	defer ll.mutex.RUnlock()

	var fields []string
	for loggerName, lvl := range ll.specs {
		fields = append(fields, fmt.Sprintf("%s=%s", loggerName, lvl))
	}
	sort.Strings(fields)
	fields = append(fields, ll.defaultLevel.String())
	return strings.Join(fields, ":")
}

// Enabled 只要给定的日志等级大于或等于 LoggerLevels 的 minLevel，就会返回 true。
func (ll *LoggerLevels) Enabled(lvl zapcore.Level) bool {
	ll.mutex.RLock()
	enabled := ll.minLevel.Enabled(lvl)
	ll.mutex.RUnlock()
	return enabled
}

func (ll *LoggerLevels) cachedLevel(loggerName string) (lvl zapcore.Level, ok bool) {
	ll.mutex.RLock()
	lvl, ok = ll.levelCache[loggerName]
	ll.mutex.RUnlock()
	return
}

func (ll *LoggerLevels) calculateLevel(loggerName string) zapcore.Level {
	candidate := loggerName + "."
	for {
		if lvl, ok := ll.specs[candidate]; ok {
			return lvl
		}
		idx := strings.LastIndex(candidate, ".")
		if idx < 0 {
			return ll.defaultLevel
		}
		candidate = candidate[:idx]
	}
}

func isValidLoggerName(loggerName string) bool {
	return loggerNameRegexp.MatchString(loggerName)
}