// Copyright 2018 The Vogo Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package clog

import (
	"context"
	"io"
)

var logger = NewClog()

//Info log
func Info(ctx context.Context, format string, v ...interface{}) {
	logger.Info(ctx, format, v...)
}

//Debug log
func Debug(ctx context.Context, format string, v ...interface{}) {
	logger.Debug(ctx, format, v...)
}

//Warn log
func Warn(ctx context.Context, format string, v ...interface{}) {
	logger.Warn(ctx, format, v...)
}

//Error log
func Error(ctx context.Context, format string, v ...interface{}) {
	logger.Error(ctx, format, v...)
}

//Fatal log
func Fatal(ctx context.Context, format string, v ...interface{}) {
	logger.Fatal(ctx, format, v...)
}

//Log write log
func Log(level Level, ctxInfo, output string) {
	logger.Log(level, ctxInfo, output)
}

//Logf write format log
func Logf(level Level, ctxInfo, format string, args ...interface{}) {
	logger.Logf(level, ctxInfo, format, args...)
}

//SetOutput set log output
func SetOutput(output io.Writer) {
	logger.SetOutput(output)
}

//SetLevelByString set log level by string
func SetLevelByString(level string) {
	logger.SetLevelByString(level)
	globalLogLevel = StringToLevel(level)
}

//SetContextFommatter set context formmater
func SetContextFommatter(ctxFmt ContextFormatter) {
	logger.ctxFmt = ctxFmt
}
