// Copyright 2018 The Vogo Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package clog

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync/atomic"
)

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
		ctxFmt: DefaultContextFormatter,
	}
}

// clog level
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

//Log write log to output,without checking level
func (clog *Clog) Log(level Level, ctxInfo, output string) {
	fmt.Fprintln(clog.output, formatOutput(level, clog.hideCallstack, ctxInfo, output, 3))
}

//Logf write format log to output,without checking level
func (clog *Clog) Logf(level Level, ctxInfo, format string, args ...interface{}) {
	fmt.Fprintln(clog.output, formatOutput(level, clog.hideCallstack, ctxInfo, fmt.Sprintf(format, args...), 3))
}

func (clog *Clog) levelContextLog(ctx context.Context, level Level, format string, args ...interface{}) {
	if clog.level() < level {
		return
	}

	ctxInfo := clog.ctxFmt(ctx)
	fmt.Fprintln(clog.output, formatOutput(level, clog.hideCallstack, ctxInfo, fmt.Sprintf(format, args...), 4))
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
