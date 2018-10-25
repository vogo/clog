// Copyright 2018 The Vogo Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package clog

//Level log level
type Level uint32

const (
	//FatalLevel fatal level
	FatalLevel Level = iota
	//ErrorLevel error level
	ErrorLevel
	//WarnLevel warn level
	WarnLevel
	//InfoLevel info level
	InfoLevel
	//DebugLevel debug level
	DebugLevel
)

var globalLogLevel = InfoLevel

//GlobalLevel global log level
func GlobalLevel() Level {
	return globalLogLevel
}

//DebugEnabled whether debug enabled
func DebugEnabled() bool {
	return globalLogLevel >= DebugLevel
}

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	}

	return "unknown"
}

//StringToLevel parse string level
func StringToLevel(level string) Level {
	switch level {
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warn", "warning":
		return WarnLevel
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	}
	return InfoLevel
}
