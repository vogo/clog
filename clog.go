// Copyright 2018 The Vogo Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package clog

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

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

//ContextFormatter format a context to string
type ContextFormatter func(ctx context.Context) string

//Clog struct
type Clog struct {
	Level         Level
	output        io.Writer
	hideCallstack bool
	ctxFmt        ContextFormatter
}

//NewClog create a new clog
func NewClog() *Clog {
	return &Clog{
		Level:  globalLogLevel,
		output: os.Stdout,
		ctxFmt: func(ctx context.Context) string {
			return "-"
		},
	}
}
func (clog *Clog) level() Level {
	return Level(atomic.LoadUint32((*uint32)(&clog.Level)))
}

//SetLevel set level
func (clog *Clog) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&clog.Level), uint32(level))
}

//SetLevelByString set level by string
func (clog *Clog) SetLevelByString(level string) {
	clog.SetLevel(StringToLevel(level))
}

//SetContextFommatter set context formmater
func (clog *Clog) SetContextFommatter(ctxFmt ContextFormatter) {
	clog.ctxFmt = ctxFmt
}

var replacer = strings.NewReplacer("\r", "\\r", "\n", "\\n")

// formatOutput format output
func (clog *Clog) formatOutput(level Level, ctxInfo, output string, callerDepth int) string {
	now := time.Now().Format("20060102 15:04:05.99999")

	output = replacer.Replace(output)

	if clog.hideCallstack {
		return fmt.Sprintf("%-25s %-5s [%s] %s",
			now, strings.ToUpper(level.String()), ctxInfo, output)
	}
	_, file, line, ok := runtime.Caller(callerDepth)
	if !ok {
		file = "???"
		line = 0
	}
	// short file name
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			file = file[i+1:]
			break
		}
	}
	return fmt.Sprintf("%-25s %-5s [%s] %s (%s:%d)",
		now, strings.ToUpper(level.String()), ctxInfo, output, file, line)
}

//Log write log to output,without checking level
func (clog *Clog) Log(level Level, ctxInfo, output string) {
	fmt.Fprintln(clog.output, clog.formatOutput(level, ctxInfo, output, 3))
}

//Logf write format log to output,without checking level
func (clog *Clog) Logf(level Level, ctxInfo, format string, args ...interface{}) {
	fmt.Fprintln(clog.output, clog.formatOutput(level, ctxInfo, fmt.Sprintf(format, args...), 3))
}

func (clog *Clog) levelContextLog(ctx context.Context, level Level, format string, args ...interface{}) {
	if clog.level() < level {
		return
	}

	ctxInfo := clog.ctxFmt(ctx)
	fmt.Fprintln(clog.output, clog.formatOutput(level, ctxInfo, fmt.Sprintf(format, args...), 4))
}

//Debug log
func (clog *Clog) Debug(ctx context.Context, format string, args ...interface{}) {
	clog.levelContextLog(ctx, DebugLevel, format, args...)
}

//Info log
func (clog *Clog) Info(ctx context.Context, format string, args ...interface{}) {
	clog.levelContextLog(ctx, InfoLevel, format, args...)
}

//Warn log
func (clog *Clog) Warn(ctx context.Context, format string, args ...interface{}) {
	clog.levelContextLog(ctx, WarnLevel, format, args...)
}

//Error log
func (clog *Clog) Error(ctx context.Context, format string, args ...interface{}) {
	clog.levelContextLog(ctx, ErrorLevel, format, args...)
}

//Fatal log
func (clog *Clog) Fatal(ctx context.Context, format string, args ...interface{}) {
	clog.levelContextLog(ctx, FatalLevel, format, args...)
}

//SetOutput set log output
func (clog *Clog) SetOutput(output io.Writer) *Clog {
	clog.output = output
	return clog
}

// HideCallstack whether hiden call stack
func (clog *Clog) HideCallstack() *Clog {
	clog.hideCallstack = true
	return clog
}

var clog = NewClog()

//Info log
func Info(ctx context.Context, format string, v ...interface{}) {
	clog.Info(ctx, format, v...)
}

//Debug log
func Debug(ctx context.Context, format string, v ...interface{}) {
	clog.Debug(ctx, format, v...)
}

//Warn log
func Warn(ctx context.Context, format string, v ...interface{}) {
	clog.Warn(ctx, format, v...)
}

//Error log
func Error(ctx context.Context, format string, v ...interface{}) {
	clog.Error(ctx, format, v...)
}

//Fatal log
func Fatal(ctx context.Context, format string, v ...interface{}) {
	clog.Fatal(ctx, format, v...)
}

//Log write log
func Log(level Level, ctxInfo, output string) {
	clog.Log(level, ctxInfo, output)
}

//Logf write format log
func Logf(level Level, ctxInfo, format string, args ...interface{}) {
	clog.Logf(level, ctxInfo, format, args...)
}

//SetOutput set log output
func SetOutput(output io.Writer) {
	clog.SetOutput(output)
}

//SetLevelByString set log level by string
func SetLevelByString(level string) {
	clog.SetLevelByString(level)
	globalLogLevel = StringToLevel(level)
}

//SetContextFommatter set context formmater
func SetContextFommatter(ctxFmt ContextFormatter) {
	clog.ctxFmt = ctxFmt
}
